package auth

import (
	"github.com/miguelmartinez624/mmarket/modules/dto"
	"github.com/miguelmartinez624/mmarket/nodos"
)

func (m *Module) Join(red nodos.NeuralRed) {
	m.OnAccountCreated = func(ev *dto.AccountRegisterEventData) {
		redEvent := nodos.Event{
			Name: nodos.ACCOUNT_CREATED,
			Data: ev,
		}

		red.Emit("authentication", redEvent)

	}

}

// this one doesn't need to listen to any channel.
func (a *Module) ListenEvents(net nodos.NeuralRed) {

}
