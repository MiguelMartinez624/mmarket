package accounts

type Encrypter interface {
	HashPassword(password string) (hash string, err error)

	ValidateHash(original string, underTest string) (success bool, err error)

	GenerateValidationHash(accountID string, seed string) (hash string, err error)
}
