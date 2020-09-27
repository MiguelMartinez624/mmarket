package events

import "github.com/miguelmartinez624/mmarket/modules/authentication/core/accounts"

type AccountCreatedEventData struct {
	Keys     accounts.NewAccountKeys
	Resource interface{} // the profile itself
}
