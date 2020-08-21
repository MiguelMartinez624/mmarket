package auth

import (
	"context"
	"fmt"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/records"
	"github.com/miguelmartinez624/mmarket/modules/nodos"
	"log"
)

type Module struct {
	AccountsService    accounts.Service
	LoginAttempService *records.Service

	tokenManager TokenManager
	notify       nodos.EventHandler
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
func (m *Module) RegisterAccounts(ctx context.Context, register *dto.RegisterUser) (success bool, err error) {

	keys, err := m.AccountsService.CreateAccount(ctx, register.Username, register.Password)
	if err != nil {
		return false, err
	}

	log.Printf("enviar a comunication %v")

	// Sent the Resource and the ID that is under the account for that resource
	evData := dto.AccountRegisterEventData{
		ResourceID: keys.ResourceID,
		Resource:   register.Resource,
	}

	//Create and sent the event.
	ev := nodos.Event{Name: ACCOUNT_CREATED, Data: evData}
	m.notify(ev)

	return true, nil
}

func (m *Module) Authenticate(ctx context.Context, loginAccount *dto.LoginAccount) (token string, err error) {
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
	ev := nodos.Event{Name: ACCOUNT_EMAIL_VERIFIED, Data: account.Username}
	m.notify(ev)

	return true, nil
}

func (m *Module) ValidateToken(ctx context.Context, token string) (claims *TokenClaims, err error) {
	return m.tokenManager.ValidateToken(token)
}

func (m *Module) SetNotificationHandler(handler nodos.EventHandler) {
	m.notify = handler
}

func (m *Module) ListenEvents(net chan nodos.Event) {
	for ev := range net {
		fmt.Println(ev)
	}
}
