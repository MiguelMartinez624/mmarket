package facade

import (
	"context"
	"fmt"

	"github.com/gompany/core/authentication/domains/accounts"
	"github.com/gompany/core/authentication/domains/records"
	"github.com/gompany/core/authentication/dto"
	"github.com/gompany/core/authentication/external"
)

type Authentication struct {
	AccountsService    *accounts.Service
	LoginAttempService *records.Service

	profileModule external.ProfileModule
}

func NewAuthentication(accountRepository accounts.Repository, encrypter accounts.Encrypter) *Authentication {
	credService := accounts.NewService(accountRepository, encrypter)

	auth := Authentication{AccountsService: credService}
	return &auth
}

//Add a account
func (m *Authentication) RegisterAccounts(ctx context.Context, register *dto.RegisterUser) (success bool, err error) {
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

func (m *Authentication) Authenticate(ctx context.Context, loginAccount *dto.LoginAccount) (resource interface{}, err error) {
	account, err := m.AccountsService.Authenticate(ctx, loginAccount.Username, loginAccount.Password)

	resource, err = m.profileModule.GetProfileByAccountID(account.ID)

	return
}
