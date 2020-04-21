package users

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/users/domains/profile"
)

// Module for the users profile domian administration
type Module struct {
	profileService *profile.Service
}

func BuildModule(profileStore profile.Store) *Module {
	service := profile.NewService(profileStore)
	m := Module{profileService: service}
	return &m
}

func (m *Module) CreateNewUserProfile(ctx context.Context, p *profile.Profile) (ID string, err error) {
	return m.profileService.CreateProfile(ctx, p)

}
