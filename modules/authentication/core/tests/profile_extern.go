package tests

import (
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
)

//MuckRepository its a muck depency for testing
type MuckProfileModule struct {
	CreateProfileFunc         func(profile *dto.Profile) (success bool, err error)
	GetProfileByAccountIDFunc func(accID string) (profile *dto.Profile, err error)
}

func (r *MuckProfileModule) GetProfileByAccountID(accID string) (profile *dto.Profile, err error) {
	return r.GetProfileByAccountIDFunc(accID)
}

func (r *MuckProfileModule) CreateProfile(profile *dto.Profile) (success bool, err error) {
	return r.CreateProfileFunc(profile)
}

//Success cases
func CreateProfileSuccess(profile *dto.Profile) (success bool, err error) {
	return true, nil
}

//Success cases
func GetProfileByAccountIDSuccess(accID string) (profile *dto.Profile, err error) {
	return &dto.Profile{AccountID: accID, FirstName: "miguel", LastName: "martinez", ID: "demo"}, nil
}
