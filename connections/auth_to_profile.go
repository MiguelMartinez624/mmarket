package connections

import (
	authDto "github.com/gompany/core/authentication/dto"
	"github.com/miguelmartinez624/mmarket/modules/users"
)

func AuthToProfileConnection(u *users.Module) *APC {
	return &APC{m: u}
}

type APC struct {
	m *users.Module
}

func (c *APC) CreateProfile(profile *authDto.Profile) (success bool, err error) {
	return
}

func (c *APC) GetProfileByAccountID(accID string) (account *authDto.Profile, err error) {
	return
}
