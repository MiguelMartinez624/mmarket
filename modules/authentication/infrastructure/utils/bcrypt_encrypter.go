package utils

import "golang.org/x/crypto/bcrypt"

type BcryptEncripter struct{}

func (b BcryptEncripter) HashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b BcryptEncripter) ValidateHash(original string, underTest string) (success bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(original), []byte(underTest))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b BcryptEncripter) GenerateValidationHash(key string, seed string) (hash string, err error) {
	return "hash", nil
}
