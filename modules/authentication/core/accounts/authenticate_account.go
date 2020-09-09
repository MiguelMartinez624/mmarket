package accounts

import (
	"context"
	cErr "github.com/miguelmartinez624/mmarket/nodos/errors"
)


func (cs *DefaultService) Authenticate(ctx context.Context, username string, password string) (account *Account, err error) {

	account, err = cs.accountRepository.GetAccountsByUserName(ctx, username)
	if err != nil {
		switch err.(type) {
		case cErr.DontExist:
			return nil, InvalidAccountsError

		default:
			return nil, err
		}
	}

	//validate password
	if success, err := cs.encrypter.ValidateHash(account.Password, password); err != nil {
		return nil, err
	} else if !success {
		return nil, InvalidAccountsError
	}

	//Check account current Status and return the corresponde payload
	// according the te account's status
	switch account.Status {
	case Blocked:
		return nil, AccountBlockedError
	case Unverified:
		return nil, UnverifiedAccountError{}

	// if active or default return the accont with no error
	case Active:
	default:
		return account, nil
	}

	return
}
