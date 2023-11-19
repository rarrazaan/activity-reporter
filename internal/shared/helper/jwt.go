package helper

import (
	"mini-socmed/internal/constant"
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/shared/dto"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type (
	AccessJWTClaim struct {
		jwt.RegisteredClaims
		UserId    int64              `json:"user_id"`
		TokenType constant.TokenType `json:"token_type"`
	}
	RefreshJWTClaim struct {
		jwt.RegisteredClaims
		TokenType constant.TokenType `json:"token_type"`
	}
)

type JwtTokenizer interface {
	GenerateAccessToken(user dto.UserTokenDTO, config dependency.Config) (*string, error)
	GenerateRefreshToken(config dependency.Config) (*string, error)
	ValidateRefreshToken(refreshToken string, config dependency.Config) (*jwt.Token, error)
}

type jwtTokenizer struct{}

func NewJwtTokenizer() JwtTokenizer {
	return &jwtTokenizer{}
}

func (c AccessJWTClaim) Valid() error {
	now := time.Now()
	if !c.VerifyExpiresAt(now, true) {
		return ErrAccessTokenExpired
	}

	if c.TokenType != constant.AccessTokenType {
		return ErrInvalidTokenType
	}

	return nil
}

func (c RefreshJWTClaim) Valid() error {
	now := time.Now()
	if !c.VerifyExpiresAt(now, true) {
		return ErrRefreshTokenExpired
	}

	if c.TokenType != constant.RefreshTokenType {
		return ErrInvalidTokenType
	}

	return nil
}

func (j *jwtTokenizer) GenerateAccessToken(user dto.UserTokenDTO, config dependency.Config) (*string, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(config.Jwt.AccessTokenExpiration))
	now := time.Now()

	registeredClaims := jwt.RegisteredClaims{
		Issuer:    config.App.AppName,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	claims := AccessJWTClaim{
		RegisteredClaims: registeredClaims,
		UserId:           user.ID,
		TokenType:        constant.AccessTokenType,
	}

	accessToken := jwt.NewWithClaims(constant.JWTSigningMethod, claims)
	t, err := accessToken.SignedString([]byte(config.Jwt.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (j *jwtTokenizer) GenerateRefreshToken(config dependency.Config) (*string, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(config.Jwt.RefreshTokenExpiration))
	now := time.Now()

	registeredClaims := jwt.RegisteredClaims{
		Issuer:    config.App.AppName,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	claims := RefreshJWTClaim{
		RegisteredClaims: registeredClaims,
		TokenType:        constant.RefreshTokenType,
	}

	refreshToken := jwt.NewWithClaims(constant.JWTSigningMethod, claims)

	t, err := refreshToken.SignedString([]byte(config.Jwt.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func ValidateAccessToken(generateToken string, config dependency.Config) (*jwt.Token, error) {
	computeFunction := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidJWTToken
		}

		return []byte(config.Jwt.JWTSecret), nil
	}

	token, err := jwt.ParseWithClaims(generateToken, new(AccessJWTClaim), computeFunction)
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e, ok := e.Inner.(*AppError); ok {
				return nil, e
			}

			return nil, err
		}
	}

	return token, nil
}

func (j *jwtTokenizer) ValidateRefreshToken(refreshToken string, config dependency.Config) (*jwt.Token, error) {
	var computeFunction jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidJWTToken
		}

		return []byte(config.Jwt.JWTSecret), nil
	}

	claim := new(RefreshJWTClaim)
	token, err := jwt.ParseWithClaims(refreshToken, claim, computeFunction)
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e, ok := e.Inner.(*AppError); ok {
				return nil, e
			}

			return nil, err
		}
	}

	return token, nil
}

func ParseAccessTokenClaim(accessToken string, config dependency.Config) (*AccessJWTClaim, error) {
	token, _ := ValidateAccessToken(accessToken, config)
	if t, ok := token.Claims.(*AccessJWTClaim); ok {
		return t, nil
	}
	return nil, ErrInvalidJWTToken
}
