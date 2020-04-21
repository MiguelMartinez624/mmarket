package profiles

import (
	"context"
)

type Service struct {
	profileStore Store
	validator    Validator
}

func NewService(profileStore Store) *Service {
	return &Service{profileStore: profileStore}
}

func (s *Service) CreateProfile(ctx context.Context, profile *Profile) (ID string, err error) {
	err = s.validator.ValidateProfile(profile)
	if err != nil {
		return "", err
	}

	//We set as main a unferevied the first contact info
	profile.Contacts[0].ItsMain = true
	profile.Contacts[0].ItsVerified = false

	ID, err = s.profileStore.StoreProfile(ctx, profile)
	if err != nil {
		return "", err
	}

	return
}

func (s *Service) GetProfileByAccountID(ctx context.Context, accountID string) (p *Profile, err error) {

	if accountID == "" {
		return nil, MissingParamError{Param: "accountID"}
	}

	p, err = s.profileStore.FindProfileByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return p, err
}
