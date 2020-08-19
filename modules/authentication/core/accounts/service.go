package accounts



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

