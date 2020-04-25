package utils

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/accounts"
)

var secretSign = []byte("secretclae")

type TokenClaims struct {
	AccountID string `json:"account_id"`
}

func (t TokenClaims) Valid() error {
	if t.AccountID == "" {
		return errors.New("missing account ID")
	}
	return nil
}

type JWTTokenManager struct{}

func (t JWTTokenManager) GenerateToken(account *accounts.Account) (token string, err error) {

	claims := TokenClaims{
		AccountID: account.ID,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString(secretSign)
	if err != nil {
		return "nil", fmt.Errorf("Error signining")
	}
	return token, nil
}

// ValidateToken use to validate json token ang get claims data
func (t JWTTokenManager) ValidateToken(tokenString string) (accountId string, err error) {

	// Parse the token
	claims := &TokenClaims{}
	tk, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify

		return secretSign, nil
	})
	if err != nil {
		return "", err
	}

	err = tk.Claims.Valid()
	if err != nil {
		return "", err
	}

	accountId = claims.AccountID
	return accountId, nil
}
