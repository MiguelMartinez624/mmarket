package users

import (
	"context"
	"github.com/miguelmartinez624/mmarket/modules/users/core/profiles"
	"github.com/miguelmartinez624/mmarket/nodos"
	"log"
	"time"
)

func (m *Module) Join(net *nodos.NeuralRed) {

	if authCon := net.Connections["authentication"]; authCon != nil {
		for ev := range authCon {
			switch ev.Name {
			case nodos.ACCOUNT_CREATED:
				log.Printf("INCOMING :: %s", ev)
				// on fail case
				var profile profiles.Profile
				if err := ev.GetData(&profile); err != nil {
					log.Println(err)
					return
				}

				//succeed case
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				m.CreateNewUserProfile(ctx, &profile)
				break
			default:
				log.Println("Unhandled event.")
				log.Println(ev)
			}
		}
	}


}

