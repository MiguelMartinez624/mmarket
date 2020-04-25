package auth

import "github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/accounts"

type TokenManager interface {
	GenerateToken(account *accounts.Account) (token string, err error)
}
