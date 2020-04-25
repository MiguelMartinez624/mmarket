package connections

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/dto"
	users "github.com/miguelmartinez624/mmarket/modules/users/core"
)

type StoreToProfile struct {
	m *users.Module
}

func StoreToProfileConnection(u *users.Module) *StoreToProfile {
	return &StoreToProfile{m: u}
}
func (c *StoreToProfile) GetProfileByID(ctx context.Context, ID string) (profile *dto.Profile, err error) {

	pf, err := c.m.GetProfilebyID(ctx, ID)
	if err != nil {
		return nil, err
	}
	profile = &dto.Profile{ID: pf.ID}
	return profile, nil
}
