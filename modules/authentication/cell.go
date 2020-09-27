package auth

import (
	"github.com/miguelmartinez624/mmarket/events"
	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"
	"github.com/miguelmartinez624/mmarket/nodos"
	"log"
)

type AuthCell struct {
	module *auth.Module
}

func (c *AuthCell) Join(red *nodos.NeuralRed) {
	c.module.OnAccountCreated = func(ev *accounts.NewAccountKeys,resource interface {},err error) {

		//TODO probably move this objets to the module as result object
		// or make this global events and used only by the cells implementation
		data := events.AccountCreatedEventData{
			Keys: *ev,
			Resource: resource,
		}

		redEvent := nodos.Event{
			Name: nodos.ACCOUNT_CREATED,
			Data: data,
		}
		log.Println(redEvent)
		red.Emit("authentication", redEvent)

	}

}
