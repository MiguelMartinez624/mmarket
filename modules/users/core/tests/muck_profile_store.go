package tests

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/users/core/domains/profiles"
)

var DATABASE map[string]*profiles.Profile = make(map[string]*profiles.Profile)

type MuckProfileStore struct{}

func (s *MuckProfileStore) StoreProfile(ctx context.Context, profile *profiles.Profile) (ID string, err error) {
	profile.ID = string(len(DATABASE))
	DATABASE[profile.ID] = profile
	return profile.ID, nil
}

func (s *MuckProfileStore) FindProfileByID(ctx context.Context, ID string) (profile *profiles.Profile, err error) {
	for _, p := range DATABASE {
		if p.ID == ID {
			profile = p
		}
	}
	if profile == nil {
		return nil, profiles.ProfileDontFoundError{}
	}

	return
}

func (s *MuckProfileStore) FindProfileByAccountID(ctx context.Context, accountID string) (profile *profiles.Profile, err error) {
	for _, p := range DATABASE {
		if p.ID == accountID {
			profile = p
		}
	}
	if profile == nil {
		return nil, profiles.ProfileDontFoundError{}
	}
	return
}

func (s *MuckProfileStore) UpdateProfile(ctx context.Context, ID string, profile *profiles.Profile) (success bool, err error) {
	profile, _ = s.FindProfileByID(ctx, ID)

	profile.FirstName = profile.FirstName
	profile.LastName = profile.LastName
	return true, nil
}
