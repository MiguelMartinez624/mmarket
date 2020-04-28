package connections

import (
	"context"
	"fmt"

	authDto "github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
	users "github.com/miguelmartinez624/mmarket/modules/users/core"
	pd "github.com/miguelmartinez624/mmarket/modules/users/core/domains/profiles"
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
		Contacts: []pd.ContactInfo{
			{Value: profile.Email, Channel: pd.Email},
		},
		Roles: []string{profile.Role},
	}
	_, err = c.m.CreateNewUserProfile(ctx, &p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *APC) GetProfileByAccountID(accID string) (profile *authDto.Profile, err error) {
	ctx := context.TODO()
	p, err := c.m.GetAccountProfile(ctx, accID)
	if err != nil {
		return nil, err
	}
	profile = &authDto.Profile{
		AccountID: p.AccountID,
		Email:     p.Contacts[0].Value,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		ID:        p.ID}

	return profile, nil
}

func (c *APC) ValidateEmail(accID string) (success bool, err error) {
	ctx := context.TODO()

	return c.m.ValidateContact(ctx, accID)
}
