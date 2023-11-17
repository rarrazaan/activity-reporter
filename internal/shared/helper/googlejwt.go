package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	jwt.RegisteredClaims
}

type googleJWT struct{}

type GoogleJWT interface {
}

func NewGoogleJWT() *googleJWT {
	return &googleJWT{}
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}

	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}

	return key, nil
}

func (g *googleJWT) ValidateGoogleJWT(tokenString string) (GoogleClaims, error) {
	claimsStruct := GoogleClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}

			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		return GoogleClaims{}, ErrInvalidJWTToken
	}
	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return GoogleClaims{}, ErrInvalidJWTToken
	}
	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {

		return GoogleClaims{}, ErrInvalidJWTToken

	}

	return *claims, nil

}
