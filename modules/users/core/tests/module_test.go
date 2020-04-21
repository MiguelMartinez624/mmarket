package tests

import (
	"testing"

	users "github.com/miguelmartinez624/mmarket/modules/users/core"
	"github.com/miguelmartinez624/mmarket/modules/users/core/domains/profiles"
	"golang.org/x/net/context"
)

var store = &MuckProfileStore{}
var ctx = context.TODO()

type ModuleTestFunction func(s *users.Module, t *testing.T)

type TestCase struct {
	Name     string
	Callback ModuleTestFunction
}

func ModuleSuite(s *users.Module, t *testing.T) {
	tc := []TestCase{
		{Name: "Attemp create with no accountID", Callback: AttempToCreateWithMissingAccount},
		{Name: "Create profile succed", Callback: CreateProfileSucceded},
	}

	for _, tCase := range tc {
		t.Run(tCase.Name, func(t *testing.T) {
			tCase.Callback(s, t)
		})
	}
}

func AttempToCreateWithMissingAccount(s *users.Module, t *testing.T) {
	badProfile := profiles.Profile{}

	_, err := s.CreateNewUserProfile(ctx, &badProfile)
	if err == nil {
		t.Error("It should have returned a error")
	}

	switch te := err.(type) {
	case profiles.MissingAccountIDError:
		return
	default:
		t.Errorf("expected [%v] to be [MissingAccountIDError]", te)
	}
}

func CreateProfileSucceded(s *users.Module, t *testing.T) {
	profile := profiles.Profile{
		AccountID: "1",
		Contacts:  []profiles.ContactInfo{},
		FirstName: "Miguel",
		LastName:  "Olivarez"}

	_, err := s.CreateNewUserProfile(ctx, &profile)
	if err != nil {
		t.Error("It should have returned a error")
	}

}

func TestModule(t *testing.T) {

	module := users.BuildModule(store)
	// module.ConnectToProfiles(profileMuck)
	ModuleSuite(module, t)
}
