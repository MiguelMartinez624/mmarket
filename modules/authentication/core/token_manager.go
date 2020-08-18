package auth

import (
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"
)

type TokenManager interface {
	GenerateToken(account *accounts.Account, profileID string) (token string, err error)
	ValidateToken(token string) (claims *TokenClaims, err error)
}
