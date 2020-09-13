package accounts

import (
	"context"
	"errors"
)

// ValidateAccountWithHash validates an account using a hash
func (cs *DefaultService) ValidateAccountWithHash(ctx context.Context, hash string) (acc *Account, err error) {

	acc, err = cs.accountRepository.GetAccountsByValidationHash(ctx, hash)
	if err != nil {
		switch err.(type) {
		case AccountDontExist:
			return nil, InvalidVerificationCodeError
		default:
			return nil, err
		}
	}

	switch acc.Status {
	case Blocked:
		return nil, AccountBlockedError
	case Active:
		return nil, AccountAlreadyVerifiedError
	}

	//Declare data to be updated
	updateData := &Account{Status: Active}

	//Need yo Update
	success, err := cs.accountRepository.UpdateAccount(ctx, acc.ID, updateData)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, errors.New("Cound update,")
	}

	// remove password
	acc.Password = ""
	return acc, nil
}
