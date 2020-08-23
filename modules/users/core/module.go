package users

import (
	"context"
	"github.com/miguelmartinez624/mmarket/nodos"
	"log"
	"time"

	"github.com/miguelmartinez624/mmarket/modules/users/core/profiles"
)

// Module for the users profile domian administration
type Module struct {
	profileService *profiles.Service
	notify         nodos.EventHandler
}

func (m *Module) SetNotificationHandler(handler nodos.EventHandler) {
	m.notify = handler
}

func (m *Module) ListenEvents(net chan nodos.Event) {
	for ev := range net {
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

func BuildModule(profileStore profiles.Store) *Module {
	service := profiles.NewService(profileStore)
	m := Module{profileService: service}
	return &m
}

func (m *Module) CreateNewUserProfile(ctx context.Context, p *profiles.Profile) (ID string, err error) {
	ID, err = m.profileService.CreateProfile(ctx, p)
	if err != nil {
		m.notify(nodos.Event{Name: nodos.PROFILE_ERROR, Data: err})
	}
	return
}

func (m *Module) GetAccountProfile(ctx context.Context, accountID string) (ID *profiles.Profile, err error) {
	return m.profileService.GetProfileByAccountID(ctx, accountID)
}

func (m *Module) ValidateContact(ctx context.Context, accountID string) (sucess bool, err error) {
	return m.profileService.ValidateMainContactInfo(ctx, accountID, true)
}

func (m *Module) GetProfilebyID(ctx context.Context, profileId string) (profile *profiles.Profile, err error) {
	return m.profileService.GetProfileByID(ctx, profileId)
}

func (m *Module) UpdateProfile(ctx context.Context, ID string, update profiles.Profile) (ok bool, err error) {
	return m.profileService.UpdateProfile(ctx, ID, update)
}
