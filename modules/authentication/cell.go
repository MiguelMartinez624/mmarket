package auth

import (
	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
	"github.com/miguelmartinez624/mmarket/modules/dto"
	"github.com/miguelmartinez624/mmarket/nodos"
)

type AuthCell struct {
	module *auth.Module
}

func (c *AuthCell) Join(red *nodos.NeuralRed) {
	c.module.OnAccountCreated = func(ev *dto.AccountRegisterEventData) {
		redEvent := nodos.Event{
			Name: nodos.ACCOUNT_CREATED,
			Data: ev,
		}

		red.Emit("authentication", redEvent)

	}

}
