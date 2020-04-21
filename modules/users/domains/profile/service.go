package profile

import (
	"context"
)

type Service struct {
	profileStore Store
	validator    Validator
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateProfile(ctx context.Context, profile *Profile) (ID string, err error) {
	err = s.validator.ValidateProfile(profile)
	if err != nil {
		return "", err
	}

	ID, err = s.profileStore.StoreProfile(ctx, profile)
	if err != nil {
		return "", err
	}

	return
}
