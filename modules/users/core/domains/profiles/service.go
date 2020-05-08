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

func (s *Service) GetProfileByID(ctx context.Context, profileID string) (p *Profile, err error) {

	if profileID == "" {
		return nil, MissingParamError{Param: "profileID"}
	}

	p, err = s.profileStore.FindProfileByID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (s *Service) ValidateMainContactInfo(ctx context.Context, accountID string, valid bool) (success bool, err error) {
	p, err := s.profileStore.FindProfileByAccountID(ctx, accountID)
	if err != nil {
		return false, err
	}

	for _, cont := range p.Contacts {
		if cont.ItsMain == true {
			cont.ItsVerified = valid
			break
		}
	}

	// update the profile here.
	success, err = s.profileStore.UpdateProfile(ctx, p.ID, p)
	if err != nil {
		return false, err
	}
	return
}

func (s *Service) UpdateProfile(ctx context.Context, profileID string, changes Profile) (ok bool, err error) {
	if profileID == "" {
		return false, MissingParamError{Param: "profileID"}
	}

	_, err = s.profileStore.FindProfileByID(ctx, profileID)
	if err != nil {
		return false, err
	}

	update := Profile{
		LastName:  changes.LastName,
		FirstName: changes.FirstName,
	}

	return s.profileStore.UpdateProfile(ctx, profileID, &update)
}
