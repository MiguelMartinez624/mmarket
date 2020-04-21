package connections

import (
	"context"
	"fmt"

	authDto "github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
	"github.com/miguelmartinez624/mmarket/modules/users"
	pd "github.com/miguelmartinez624/mmarket/modules/users/domains/profile"
)

func AuthToProfileConnection(u *users.Module) *APC {
	return &APC{m: u}
}

type APC struct {
	m *users.Module
}

func (c *APC) CreateProfile(profile *authDto.Profile) (success bool, err error) {
	fmt.Println(profile)
	ctx := context.TODO()
	p := pd.Profile{
		AccountID: profile.AccountID,
		LastName:  profile.LastName,
		FirstName: profile.FirstName,
	}
	_, err = c.m.CreateNewUserProfile(ctx, &p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *APC) GetProfileByAccountID(accID string) (account *authDto.Profile, err error) {
	return
}