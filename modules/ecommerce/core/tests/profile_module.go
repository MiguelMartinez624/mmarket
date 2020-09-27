package tests

import (
	"context"
	"github.com/miguelmartinez624/mmarket/modules/users/core/profiles"
)

type MuckProfileModule struct{}

func (m MuckProfileModule) GetProfileByID(ctx context.Context, ID string) (profile *profiles.Profile, err error) {

	profile = &profiles.Profile{
		ID: ID,
	}

	return profile, nil
}
