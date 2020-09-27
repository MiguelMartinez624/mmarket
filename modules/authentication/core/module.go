package auth

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/records"
)

type AccountCreateCallback func(ev *accounts.NewAccountKeys, resource interface{}, err error)
type AccountValidatedCallback func(account *accounts.Account, err error)

type Module struct {
	AccountsService    accounts.Service
	LoginAttempService *records.Service

	tokenManager TokenManager

	// Events handlers
	OnAccountCreated   AccountCreateCallback
	OnAccountValidated AccountValidatedCallback
}

func NewAuthentication(
	accountRepository accounts.Repository,
	encrypter accounts.Encrypter,
	tokenManager TokenManager) *Module {

	credService := accounts.NewService(accountRepository, encrypter)

	auth := Module{
		AccountsService: credService,
		tokenManager:    tokenManager,
	}

	return &auth
}

//RegisterAccounts register a account and sent te profile data to the profiles data.
func (m *Module) RegisterAccounts(ctx context.Context, register *RegisterUser) (success bool, err error) {

	keys, err := m.AccountsService.CreateAccount(ctx, accounts.Account{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password})
	if err != nil {
		return false, err
	}

	//Create and sent the event.
	if m.OnAccountCreated != nil {
		// Sent the Resource and the ID that is under the account for that resource
		m.OnAccountCreated(keys, register.Resource, err)
	}

	return true, nil
}

func (m *Module) Authenticate(ctx context.Context, loginAccount *LoginAccount) (token string, err error) {
	account, err := m.AccountsService.Authenticate(ctx, loginAccount.Username, loginAccount.Password)
	if err != nil {
		return "", err
	}

	token, err = m.tokenManager.GenerateToken(account, account.ResourceID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Module) ValidateAccount(ctx context.Context, hash string) (success bool, err error) {
	account, err := m.AccountsService.ValidateAccountWithHash(ctx, hash)
	if err != nil {
		return false, err
	}
	//Once the account its validated we procced to mark the email as valid
	//Create and sent the event.
	//ev := nodos.Event{Name: ACCOUNT_EMAIL_VERIFIED, Data: account.Username}
	//m.notify(ev)

	//Create and sent the event.
	if m.OnAccountValidated != nil {
		// Sent the Resource and the ID that is under the account for that resource
		m.OnAccountValidated(account, err)
	}

	return true, nil
}

func (m *Module) ValidateToken(ctx context.Context, token string) (claims *TokenClaims, err error) {
	return m.tokenManager.ValidateToken(token)
}
