package auth

import (
	"context"
	"fmt"

	"github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/accounts"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/domains/records"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/external"
)

type Module struct {
	AccountsService    *accounts.Service
	LoginAttempService *records.Service

	profileModule external.ProfileModule
}

func NewAuthentication(accountRepository accounts.Repository, encrypter accounts.Encrypter) *Module {
	credService := accounts.NewService(accountRepository, encrypter)

	auth := Module{AccountsService: credService}
	return &auth
}

//RegisterAccounts register a account and sent te profile data to the profiles data.
func (m *Module) RegisterAccounts(ctx context.Context, register *dto.RegisterUser) (success bool, err error) {
	keys, err := m.AccountsService.CreateAccount(ctx, register.Username, register.Password)
	if err != nil {
		return false, err
	}

	fmt.Printf("enviar a comunication %v", keys.VerificationHash)
	newProfile := dto.Profile{
		AccountID: keys.AccountID,
		FirstName: register.FirstName,
		LastName:  register.LastName,
	}

	success, err = m.profileModule.CreateProfile(&newProfile)
	if err != nil {
		//handle what kind of error cud happend and retry probably
		// panic(err)
	}

	return success, err
}

func (m *Module) Authenticate(ctx context.Context, loginAccount *dto.LoginAccount) (resource interface{}, err error) {
	account, err := m.AccountsService.Authenticate(ctx, loginAccount.Username, loginAccount.Password)
	if err != nil {
		return nil, err
	}
	resource, err = m.profileModule.GetProfileByAccountID(account.ID)
	if err != nil {
		return nil, err
	}

	return
}

func (m *Module) ConnectToProfiles(pm external.ProfileModule) {
	m.profileModule = pm
}
