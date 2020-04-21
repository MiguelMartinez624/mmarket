package connections

import (
	authDto "github.com/gompany/core/authentication/dto"
	authEx "github.com/gompany/core/authentication/external"
	"github.com/miguelmartinez624/mmarket/modules/users"
)

func AuthToProfileConnection(u *users.Module) authEx.ProfileModule {
	return &APC{m: u}
}

type APC struct {
	m *users.Module
}

func (c APC) CreateProfile(profile *authDto.Profile) (ID string, err error) {
	return
}

func (c APC) GetProfileByAccountID(accID string) (account *authDto.Profile, err error) {
	return
}
