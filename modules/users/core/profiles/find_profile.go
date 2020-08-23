package profiles

import "context"

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
