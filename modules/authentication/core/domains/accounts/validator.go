package accounts

// Validator it look on the fields for unvalid values al return a error if there its one
// corresponding to the error found
type Validator struct{}

func (v *Validator) ValidateAccount(cred *Account) (err error) {

	if cred.Password == "" {
		return EmptyPasswordError
	}

	if cred.Username == "" {
		return EmptyUsernameError
	}

	return nil
}
