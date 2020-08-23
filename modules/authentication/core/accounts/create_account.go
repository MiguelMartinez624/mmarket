package accounts

import (
	"context"
	"github.com/google/uuid"
	cErr "github.com/miguelmartinez624/mmarket/nodos/errors"
	"log"
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
func (cs *DefaultService) CreateAccount(ctx context.Context, acc Account) (keys *NewAccountKeys, err error) {

	log.Println("Creating account..")

	// Step 1 : validate that the account data is valid...
	if err = acc.ItsValid(); err != nil {
		return nil, err
	}

	// Step 2 : validate email avalability
	if err = cs.checkEmailAndUserAvailability(ctx, acc.Username); err != nil {
		return nil, err
	}

	// Hash and change the password to be stored as a hash
	passwordHash, err := cs.encrypter.HashPassword(acc.Password)
	if err != nil {
		return nil, err
	}

	acc.Password = passwordHash
	acc.Status = Unverified

	// Here w use the encryter to create a validation hash
	if hash, err := cs.encrypter.GenerateValidationHash("randomSeed", "SEED"); err != nil {
		return nil, err
	} else {
		acc.ValidationHash = hash
	}

	if resourceId, err := uuid.NewUUID(); err != nil {
		panic(err) // handle
	} else {
		acc.ResourceID = resourceId.String()
	}

	// persist
	ID, err := cs.accountRepository.SaveAccount(ctx, &acc)
	if err != nil {
		return nil, err
	}

	keys = &NewAccountKeys{
		AccountID:        ID,
		VerificationHash: acc.ValidationHash,
		ResourceID:       acc.ResourceID,
	}

	// OK!
	return keys, nil
}

// checkEmailAndUserAvailability Check if there its already a account to that username if its the case it will
// it will get a AlreadyExistUsernameError
func (cs *DefaultService) checkEmailAndUserAvailability(ctx context.Context, username string) error {

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
