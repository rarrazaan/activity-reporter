package helper

import (
	"activity-reporter/shared/dto"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type MyClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

type jwtTokenizer struct{}

type JwtTokenizer interface {
	GenerateToken(user dto.UserTokenDTO) (string, error)
}

func NewJwtTokenizer() *jwtTokenizer {
	return &jwtTokenizer{}
}

func (j *jwtTokenizer) GenerateToken(user dto.UserTokenDTO) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
	mySigningKey := []byte(os.Getenv("JWT_KEY"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		MyClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    os.Getenv("APP_NAME"),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			UserID: user.ID,
		})
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
	return jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidJWTToken
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
}
