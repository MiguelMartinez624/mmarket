package profile

type Repository interface {
	StoreProfile(profile *Profile) (ID string, err error)

	FindProfileByID(ID string) (profile *Profile, err error)
}
