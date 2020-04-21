package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
	"github.com/miguelmartinez624/mmarket/modules/authentication/core/dto"
)

type AuthenticationHTTP struct {
	auth *auth.Module
}

func NewAuthHTTP(auth *auth.Module) *AuthenticationHTTP {
	return &AuthenticationHTTP{auth: auth}
}

func (a *AuthenticationHTTP) Signin(w http.ResponseWriter, r *http.Request) {
	var dto dto.LoginAccount

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := a.auth.Authenticate(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&account)

}

func (a *AuthenticationHTTP) SignUp(w http.ResponseWriter, r *http.Request) {

	var dto dto.RegisterUser

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success, err := a.auth.RegisterAccounts(r.Context(), &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !success {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Check you email for confirmation code")
}
