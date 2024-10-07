package util

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySecretKeys = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	UserId string `json:"userid"`
	jwt.RegisteredClaims
}

func NewToken(id string) *claims {
	return &claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}

func (c *claims) CreateToken() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecretKeys)
}


func CheckToken(tkn string) (*string, error) {
	tokens, err := jwt.ParseWithClaims(tkn, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretKeys), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokens.Claims.(*claims)
	if !ok || !tokens.Valid {
		return nil, errors.New("token expired")
	}

	return &claims.UserId, nil
}

