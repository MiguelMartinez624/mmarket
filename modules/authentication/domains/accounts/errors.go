package accounts

import (
	"errors"
	"fmt"
)

var (
	EmptyPasswordError        error = simpleErr("Password cant be empty")
	InvalidAccountsError            = simpleErr("Invalid accounts")
	AlreadyExistUsernameError       = simpleErr("Invalid accounts")
	AccountBlockedError             = simpleErr("The accounts its currently blocked")
)

// type InvalidAccounts struct {
// }

// func (e InvalidAccounts) Error() string {
// 	return "Invalid accounts"
// }

// simpleErr create simple error with flat msg
func simpleErr(msg string) error {
	return errors.New(msg)
}

type AccountDontExist struct{}

func (e AccountDontExist) Error() string {
	return fmt.Sprintf("account doesn't exist.")
}
