package accounts

import (
	"context"
	cErr "github.com/miguelmartinez624/mmarket/nodos/errors"
)

//Authenticate method to validate and account it follows the next step to performe this action.
//1- check that the account actually exist.
//2- validate the password against the hashed password stored.
//3- checkout the state for the state of the account this may be on 3 different states.
//BLOCKED : may be for multiple reasons, it doesnt matter to the authentication process with
//the account is blocked.
//
//UNVERIFIED : the account email its not verified yet so the ownership of the email provided
//is still on a unknown state and the account cant be used.
//
//ACTIVE : this is the ideal state of the account and the only one were the account can be used.
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
