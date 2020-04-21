package accounts

import (
	"context"
)

// Service contains the logic of this domain, accounts it use its the gate
// to validate and store, search delete accounts
type Service struct {
	accountRepository Repository
	encrypter         Encrypter
	accountsValidator Validator
}

func NewService(accountRepository Repository, encrypter Encrypter) *Service {
	return &Service{
		accountRepository: accountRepository,
		encrypter:         encrypter,
	}
}

// CreateAccount create a account with the pass it username and password it response
// a keys that contains the ValidationHash required to validate the account
func (cs *Service) CreateAccount(ctx context.Context, username, password string) (keys *NewAccountKeys, err error) {

	creds := &Account{Username: username, Password: password}

	// need to validate the data that its comming here
	err = cs.accountsValidator.ValidateAccount(creds)
	if err != nil {
		return nil, err
	}

	// Check if there its already a account to that username if its the case it will
	// it will get a AlreadyExistUsernameError
	accounts, err := cs.accountRepository.GetAccountsByUserName(ctx, username)
	if err != nil {
		switch err.(type) {
		case AccountDontExist:
			break
		default:
			return nil, err
		}
	}

	// Duplicate accounts attemp error
	if accounts != nil {
		return nil, AlreadyExistUsernameError{}
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

	creds.ValidationHash = hash

	ID, err := cs.accountRepository.SaveAccount(ctx, creds)
	if err != nil {
		return nil, err
	}

	keys = &NewAccountKeys{
		AccountID:        ID,
		VerificationHash: hash,
	}

	// OK!
	return keys, nil
}

// Authenticate method to validate and account
func (cs *Service) Authenticate(ctx context.Context, username string, password string) (account *Account, err error) {

	account, err = cs.accountRepository.GetAccountsByUserName(ctx, username)
	if err != nil {
		return nil, err
	}

	//validate password
	success, err := cs.encrypter.ValidateHash(account.Password, password)
	if err != nil {
		return nil, err
	}

	if !success {
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

// ValidateAccountWithHash validates an account using a hash
func (cs *Service) ValidateAccountWithHash(ctx context.Context, hash string) (success bool, err error) {

	acc, err := cs.accountRepository.GetAccountsByValidationHash(ctx, hash)
	if err != nil {
		switch err.(type) {
		case AccountDontExist:
			return false, InvalidVerificationCodeError
		default:
			return false, err
		}
	}

	switch acc.Status {
	case Blocked:
		return false, AccountBlockedError
	case Active:
		return false, AccountAlreadyVerifiedError
	}

	acc.Status = Active

	return true, nil
}
