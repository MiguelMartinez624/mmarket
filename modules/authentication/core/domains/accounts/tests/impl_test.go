package tests

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/accounts"
)

func SaveAccount(ctx context.Context, creds *accounts.Account) (ID string, err error) {
	return "ID_TEST", nil
}
func GetAccountsByUserName(ctx context.Context, username string) (account *accounts.Account, err error) {
	return &accounts.Account{Username: username}, nil
}

type TestEncripter struct{}

func (e *TestEncripter) ValidateHash(original string, underTest string) (success bool, err error) {

	success = original == underTest

	return success, nil
}

func (e *TestEncripter) GenerateValidationHash(key string, seed string) (hast string, err error) {
	return "key-has", nil
}
func (e *TestEncripter) HashPassword(password string) (hash string, err error) {
	return "", nil
}
