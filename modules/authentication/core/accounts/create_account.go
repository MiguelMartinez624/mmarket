package accounts

import (
	"context"
	"github.com/google/uuid"
	cErr "github.com/miguelmartinez624/mmarket/modules/nodos/errors"
)

// CreateAccount create a account with the pass it username and password it response
// a keys that contains the ValidationHash required to validate the account
func (cs *Service) CreateAccount(ctx context.Context, username, password string) (keys *NewAccountKeys, err error) {

	creds := &Account{Username: username, Password: password}

	// need to validate the data that its comming here
	if err = creds.ItsValid(); err != nil {
		return nil, err
	}

	if err = cs.checkEmalAndUserAvalability(ctx, username); err != nil {
		return nil, err
	}

	// Hash and change the password to be stored as a hash
	passwordHash, err := cs.encrypter.HashPassword(creds.Password)
	if err != nil {
		return nil, err
	}

	creds.Password = passwordHash
	creds.Status = Unverified

	// Here w use the encryter to create a validation hash
	hash, err := cs.encrypter.GenerateValidationHash("randomSeed", "SEED")
	if err != nil {
		return nil, err
	}
	resourceId, err := uuid.NewUUID()
	if err != nil {
		panic(err) // handle
	}

	creds.ValidationHash = hash
	creds.ResourceID = resourceId.String()
	ID, err := cs.accountRepository.SaveAccount(ctx, creds)
	if err != nil {
		return nil, err
	}

	keys = &NewAccountKeys{
		AccountID:        ID,
		VerificationHash: hash,
		ResourceID:       creds.ResourceID,
	}

	// OK!
	return keys, nil
}

// checkEmalAndUserAvalability Check if there its already a account to that username if its the case it will
// it will get a AlreadyExistUsernameError
func (cs *Service) checkEmalAndUserAvalability(ctx context.Context, username string) error {

	accounts, err := cs.accountRepository.GetAccountsByUserName(ctx, username)
	if err != nil {
		switch err.(type) {
		case cErr.DontExist:
			break
		default:
			return err
		}
	}

	// Duplicate accounts attempt error
	if accounts != nil {
		return AlreadyExistUsernameError{}
	}
	return nil
}
