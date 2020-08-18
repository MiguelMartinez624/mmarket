package tests

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"
	"github.com/miguelmartinez624/mmarket/modules/common/repositories"
)

// In memory stored
var AccountsStore map[string]accounts.Account = make(map[string]accounts.Account)

//MuckRepository its a muck depency for testing
type MuckRepository struct {
	repositories.InMemory
	SaveAccountFunc                 func(ctx context.Context, cre *accounts.Account) (ID string, err error)
	GetAccountsByUserNameFunc       func(ctx context.Context, username string) (account *accounts.Account, err error)
	GetAccountsByValidationHashFunc func(ctx context.Context, hash string) (account *accounts.Account, err error)
}

func (r *MuckRepository) SaveAccount(ctx context.Context, acc *accounts.Account) (ID string, err error) {
	return r.Save(ctx, acc)
}

func (r *MuckRepository) GetAccountsByUserName(ctx context.Context, username string) (account *accounts.Account, err error) {
	return r.GetAccountsByUserNameFunc(ctx, username)
}
func (r *MuckRepository) GetAccountsByValidationHash(ctx context.Context, hash string) (account *accounts.Account, err error) {
	for _, acc := range AccountsStore {
		if acc.ValidationHash == hash {
			account = &acc
		}
	}
	if account == nil {
		return nil, accounts.AccountDontExist{}
	}

	return
}

func (r *MuckRepository) UpdateAccount(ctx context.Context, ID string, acc *accounts.Account) (success bool, err error) {
	AccountsStore[ID] = *acc
	return true, nil
}

//Helper funtion to modify behavior
func SaveAccount(ctx context.Context, creds *accounts.Account) (ID string, err error) {
	creds.ID = string(len(AccountsStore))
	AccountsStore[creds.ID] = *creds
	return creds.ID, nil
}
func GetAccountsByUserName(ctx context.Context, username string) (account *accounts.Account, err error) {
	for _, acc := range AccountsStore {
		if acc.Username == username {
			account = &acc
		}
	}
	if account == nil {
		return nil, accounts.AccountDontExist{}
	}

	return
}

func GetAccountsByUserNameSuccess(ctx context.Context, username string) (account *accounts.Account, err error) {
	return &accounts.Account{Username: username}, nil
}
