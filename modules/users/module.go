package users

import "github.com/gompany/core/users/profile"

// Module for the users profile domian administration
type Module struct {
	UserProfileService *profile.Service
}

func (m *Module) CreateNewUserProfile() {

}
