package utils

import "github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/accounts"

type JWTTokenManager struct{}

func (t JWTTokenManager) GenerateToken(account *accounts.Account) (token string, err error) {
	return "validToken", nil
}
