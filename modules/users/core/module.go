package users

import (
	"context"
	"github.com/miguelmartinez624/mmarket/modules/users/core/profiles"
)

// Module for the users profile domian administration
type Module struct {
	profileService *profiles.Service
}




func BuildModule(profileStore profiles.Store) *Module {
	service := profiles.NewService(profileStore)
	m := Module{profileService: service}
	return &m
}

func (m *Module) CreateNewUserProfile(ctx context.Context, p *profiles.Profile) (ID string, err error) {
	ID, err = m.profileService.CreateProfile(ctx, p)

	return
}

func (m *Module) GetAccountProfile(ctx context.Context, accountID string) (ID *profiles.Profile, err error) {
	return m.profileService.GetProfileByAccountID(ctx, accountID)
}

func (m *Module) ValidateContact(ctx context.Context, accountID string) (sucess bool, err error) {
	return m.profileService.ValidateMainContactInfo(ctx, accountID, true)
}

func (m *Module) GetProfilebyID(ctx context.Context, profileId string) (profile *profiles.Profile, err error) {
	return m.profileService.GetProfileByID(ctx, profileId)
}

func (m *Module) UpdateProfile(ctx context.Context, ID string, update profiles.Profile) (ok bool, err error) {
	return m.profileService.UpdateProfile(ctx, ID, update)
}
