package external

import "github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"

type ProfileModule interface {
	GetProfileByAccountID(accID string) (profile *dto.Profile, err error)

	CreateProfile(profile interface{}) (success bool, err error)

	ValidateEmail(accountId string) (success bool, err error)
}
