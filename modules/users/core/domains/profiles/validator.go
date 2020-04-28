package profiles

type Validator struct {
}

func (v Validator) ValidateProfile(p *Profile) error {

	if p.AccountID == "" {
		return MissingAccountIDError{}
	}

	if p.FirstName == "" {
		return MissingFirstNameError{}
	}

	if p.Roles == nil {
		return ErrMissinRole{}
	}

	if len(p.Contacts) == 0 {
		return NoContactsOnProfileError{}
	}

	return nil
}
