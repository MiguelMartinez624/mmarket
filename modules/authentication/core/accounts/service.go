package accounts

import "context"

// Service contains the logic of this domain, accounts it use its the gate
// to validate and store, search delete accounts
type Service interface {
	CreateAccount(ctx context.Context, username, password string) (keys *NewAccountKeys, err error)

	Authenticate(ctx context.Context, username string, password string) (account *Account, err error)

	ValidateAccountWithHash(ctx context.Context, hash string) (acc *Account, err error)
}

type DefaultService struct {
	accountRepository Repository
	encrypter         Encrypter
}

func NewService(accountRepository Repository, encrypter Encrypter) Service {
	return &DefaultService{
		accountRepository: accountRepository,
		encrypter:         encrypter,
	}
}
