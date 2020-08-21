package accounts

import (
	"reflect"
	"testing"
)


type createAccountTestCase struct {
	name          string
	fields        serviceDeps
	args          argsAuthenticate
	wantKeys      *NewAccountKeys
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

func TestService_CreateAccount(t *testing.T) {

	tests := []createAccountTestCase{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &Service{
				accountRepository: tt.fields.accountRepository,
				encrypter:         tt.fields.encrypter,
			}
			gotKeys, err := cs.CreateAccount(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("CreateAccount() gotKeys = %v, want %v", gotKeys, tt.wantKeys)
			}
		})
	}
}
