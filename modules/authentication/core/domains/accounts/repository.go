package accounts

import (
	"context"
)

type Repository interface {
	SaveAccount(ctx context.Context, cre *Account) (ID string, err error)

	UpdateAccount(ctx context.Context, ID string, cre *Account) (success bool, err error)

	GetAccountsByUserName(ctx context.Context, username string) (account *Account, err error)

	GetAccountsByValidationHash(ctx context.Context, hash string) (account *Account, err error)
}
