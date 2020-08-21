package accounts

import (
	"context"
	"github.com/google/uuid"
	cErr "github.com/miguelmartinez624/mmarket/modules/nodos/errors"
)

//CreateAccount take care of all steps and validations to create a account, a username and email can only belong to 1
//account and once that its taked it wont allow to create another account with it,
//
//* Step 1 : validate that the account data is valid this would be the email and the password as those are
//the only required fields for a account creation, if a username its not provided it will set the email
//as a default username that can be changed later.
//
//* Step 2 : check that the username or the email its not been used already taked by another account,
//
//* Step 3 : hash the password before stored as we never persist plain password, and set the first State
//of the account to be UNVERIFIED
//
//* Step 4 : generate a validation hash that that can be used for validate the account after its creation this
//can be send via email to the account email.
//
//* Step 5: the final step its to create a Unique ID to a resource that this account its guarding and return the
//account keys (accountID, resourceID, validationHash).
func (cs *Service) CreateAccount(ctx context.Context, username, password string) (keys *NewAccountKeys, err error) {

	creds := &Account{Username: username, Password: password}

	// need to validate the data that its coming here
	if err = creds.ItsValid(); err != nil {
		return nil, err
	}

	if err = cs.checkEmailAndUserAvalability(ctx, username); err != nil {
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

// checkEmailAndUserAvalability Check if there its already a account to that username if its the case it will
// it will get a AlreadyExistUsernameError
func (cs *Service) checkEmailAndUserAvalability(ctx context.Context, username string) error {

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
