package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"

	users "github.com/miguelmartinez624/mmarket/modules/users/core"
)

type HttpController struct {
	users *users.Module
}

func NewHttpController(users *users.Module) *HttpController {
	return &HttpController{users: users}
}

func (a *HttpController) Me(w http.ResponseWriter, r *http.Request) {
	if accountId := r.Context().Value("accountID"); accountId != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("accountID %v", accountId)
		profile, err := a.users.GetAccountProfile(r.Context(), accountId.(string))
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
