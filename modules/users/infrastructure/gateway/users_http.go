package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
	users "github.com/miguelmartinez624/mmarket/modules/users/core"
)

type HttpController struct {
	users *users.Module
}

func NewHttpController(users *users.Module) *HttpController {
	return &HttpController{users: users}
}

func (a *HttpController) Me(w http.ResponseWriter, r *http.Request) {
	if user := r.Context().Value("user"); user != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("accountID %v", user)
		profile, err := a.users.GetAccountProfile(r.Context(), user.(*auth.TokenClaims).AccountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(&profile)

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}

}
