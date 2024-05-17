package auth

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name,omitempty"`
}

var (
	key     *rsa.PrivateKey
	UserKey = "user-ctx"
)

func Configure(keyStr string) error {
	var err error

	key, err = parsePrivateKey(keyStr)
	if err != nil {
		return err
	}

	return nil
}

func parsePrivateKey(keyStr string) (*rsa.PrivateKey, error) {
	block, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		return nil, err
	}
	key, err = x509.ParsePKCS1PrivateKey(block)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}

func extractBearerToken(token string) (string, error) {
	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("unknown jwt format")
	}
	return parts[1], nil
}

func ClaimJWT(bearerToken string) (*CustomClaims, error) {
	token, err := extractBearerToken(bearerToken)
	if err != nil {
		return nil, err
	}

	claims := new(CustomClaims)
	tokenObj, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid token")
		}
		return key.Public(), nil
	})
	if err != nil {
		return nil, err
	}

	if !tokenObj.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GetUserContext(ctx context.Context) *CustomClaims {
	raw, ok := ctx.Value(UserKey).(*CustomClaims)
	if ok {
		return raw
	}

	return &CustomClaims{}
}
