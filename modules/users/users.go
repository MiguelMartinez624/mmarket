package users

import "github.com/miguelmartinez624/mmarket/modules/users/domains/profile"

// Module for the users profile domian administration
type Module struct {
	profileService *profile.Service
}

func BuildModule() *Module {
	service := profile.NewService()
	m := Module{profileService: service}
	return &m
}

func (m *Module) CreateNewUserProfile(profile profile.Profile) (ID string, err error) {
	return
}
