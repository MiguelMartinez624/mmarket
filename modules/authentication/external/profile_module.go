package external

import "github.com/gompany/core/authentication/dto"

type ProfileModule interface {
	GetProfileByAccountID(accID string) (profile dto.Profile, err error)

	CreateProfile(profile *dto.Profile) (success bool, err error)
}
