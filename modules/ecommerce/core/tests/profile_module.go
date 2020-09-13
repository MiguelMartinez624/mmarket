package tests

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/ecommerce/core/dto"
)

type MuckProfileModule struct{}

func (m MuckProfileModule) GetProfileByID(ctx context.Context, ID string) (profile *dto.Profile, err error) {

	profile = &dto.Profile{
		ID: ID,
	}

	return profile, nil
}
