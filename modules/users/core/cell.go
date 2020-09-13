package users

import (
	"context"
	"encoding/json"
	"github.com/miguelmartinez624/mmarket/modules/events"
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

				// on fail case
				 data  := ev.Data.(events.AccountCreatedEventData)
				jsonString, err := json.Marshal(data.Resource)
				if err != nil {
					panic(err)
				}
				var output profiles.Profile
				if err := json.Unmarshal(jsonString, &output); err != nil {
					panic( err)
				}

				log.Println(output)
				output.ID = data.Keys.ResourceID
				output.AccountID = data.Keys.AccountID
				output.Contacts=[]profiles.ContactInfo{
					{Value: data.Keys.Email}}
				//succeed case
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				m.CreateNewUserProfile(ctx, &output)
				break
			default:
				log.Println("Unhandled event.")
				log.Println(ev)
			}
		}
	}

}
