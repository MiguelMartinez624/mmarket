package stores

type Validator struct {
}

func (v Validator) Validate(store *Store) (err error) {
	if store.Name == "" {
		return MissinField{Field: "Name"}
	}
	if store.ProfileID == "" {
		return MissinField{Field: "ProfileID"}
	}

	return nil
}