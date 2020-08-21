package accounts

import (
	"context"
	"reflect"
	"testing"
)

type serviceDeps struct {
	accountRepository Repository
	encrypter         Encrypter
}
type argsAuthenticate struct {
	ctx      context.Context
	username string
	password string
}
type accountAuthenticateTestCase struct {
	name          string
	fields        serviceDeps
	args          argsAuthenticate
	wantAccount   *Account
	wantErr       bool
	expectedError error
}

var invalidUsernameCase = accountAuthenticateTestCase{
	name:        "Try to get account with invalid username",
	wantErr:     true,
	wantAccount: nil,
	fields: serviceDeps{
		accountRepository: MockRepository{GetByUserNameFunc: func() (account *Account, err error) {
			return nil, InvalidAccountsError
		}},
	},
	expectedError: InvalidAccountsError,
}

func TestService_Authenticate(t *testing.T) {

	tests := []accountAuthenticateTestCase{
		invalidUsernameCase,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &Service{
				accountRepository: tt.fields.accountRepository,
				encrypter:         tt.fields.encrypter,
			}
			gotAccount, err := cs.Authenticate(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("Authenticate() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}

			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.expectedError) {
					t.Errorf("Authenticate() gotError = %v, want %v", err, tt.expectedError)
				}
			}

		})
	}
}
