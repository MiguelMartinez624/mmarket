package auth

import "errors"

type TokenClaims struct {
	AccountID string `json:"account_id"`
	ProfileID string `json:"profile_id"`
}

func (t TokenClaims) Valid() error {
	if t.AccountID == "" {
		return errors.New("missing account ID")
	}
	return nil
}
