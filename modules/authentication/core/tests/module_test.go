package tests

import (
	"context"
	"testing"

	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
)

var store = &MuckRepository{
	SaveAccountFunc:           SaveAccount,
	GetAccountsByUserNameFunc: GetAccountsByUserName,
}

var profileMuck = &MuckProfileModule{
	CreateProfileFunc:         CreateProfileSuccess,
	GetProfileByAccountIDFunc: GetProfileByAccountIDSuccess,
}
var ctx = context.Background()

type ModuleTestFunction func(s *auth.Module, t *testing.T)

type TestCase struct {
	Name     string
	Callback ModuleTestFunction
}

func ModuleSuite(s *auth.Module, t *testing.T) {
	tc := []TestCase{
		{Name: "create new accounts", Callback: CreateAccountSuccessTest},
		{Name: "attemp to create already existing accounts", Callback: CreateAlreadyExistingAccountTest},
		{Name: "attemp to login on unverifid account", Callback: LoginUnVerifiedAccount},
		{Name: "verify account hash success", Callback: VerifiedAccountSuccess},
		{Name: "attemp to login on verified account", Callback: LoginVerifiedAccount},
	}

	for _, tCase := range tc {
		t.Run(tCase.Name, func(t *testing.T) {
			tCase.Callback(s, t)
		})
	}
}

func CreateAccountSuccessTest(s *auth.Module, t *testing.T) {

	rd := dto.RegisterUser{
		FirstName: "Miguel",
		LastName:  "Martiez",
		Password:  "password123",
		Username:  "username1",
		Email:     "email@test.com",
	}

	success, err := s.RegisterAccounts(ctx, &rd)
	if err != nil {
		t.Error(err)
	}

	if !success {
		t.Error("It should have succed")
	}
}

func CreateAlreadyExistingAccountTest(s *auth.Module, t *testing.T) {
	//prepare
	rd := dto.RegisterUser{
		FirstName: "Miguel",
		LastName:  "Martiez",
		Password:  "password123",
		Username:  "username1",
		Email:     "email@test.com",
	}

	_, err := s.RegisterAccounts(ctx, &rd)

	switch te := err.(type) {
	case accounts.AlreadyExistUsernameError:
		return
	case nil:
		t.Error("It should have trow and error")

	default:
		t.Errorf("expected [%v] to be [AlreadyExistUsernameError]", te)
	}
}

func LoginUnVerifiedAccount(s *auth.Module, t *testing.T) {
	creds := dto.LoginAccount{Username: "username1", Password: "password123"}
	_, err := s.Authenticate(ctx, &creds)

	switch te := err.(type) {
	case accounts.UnverifiedAccountError:
		return
	case nil:
		t.Error("It should have trow and error")

	default:
		t.Errorf("expected [%v] to be [UnverifiedAccountError]", te)
	}
}
func LoginVerifiedAccount(s *auth.Module, t *testing.T) {
	creds := dto.LoginAccount{Username: "username1", Password: "password123"}
	_, err := s.Authenticate(ctx, &creds)
	if err != nil {
		t.Errorf("expexted %v to be nil", err)
	}
}

func VerifiedAccountSuccess(s *auth.Module, t *testing.T) {
	_, err := s.ValidateAccount(ctx, "key-has")

	if err != nil {
		t.Errorf("expexted %v to be nil", err)
	}
}

func TestModule(t *testing.T) {

	encrypter := &TestEncripter{}
	tokenManager := MuckTokenManager{}
	module := auth.NewAuthentication(store, encrypter, tokenManager)
	module.ConnectToProfiles(profileMuck)
	ModuleSuite(module, t)
}
