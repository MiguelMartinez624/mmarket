package tests

type TestEncripter struct{}

func (e *TestEncripter) ValidateHash(original string, underTest string) (success bool, err error) {

	success = original == underTest

	return success, nil
}

func (e *TestEncripter) GenerateValidationHash(key string, seed string) (hast string, err error) {
	return "key-has", nil
}
func (e *TestEncripter) HashPassword(password string) (hash string, err error) {
	return password, nil
}
