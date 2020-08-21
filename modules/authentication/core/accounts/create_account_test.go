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
var emptyEmailCase = createAccountTestCase{
	name:        "Try to create email without email",
	wantErr:     true,
	wantKeys: nil,
	fields: serviceDeps{},
	expectedError: EmptyEmailError,
}

func TestService_CreateAccount(t *testing.T) {

	tests := []createAccountTestCase{
		emptyEmailCase,
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

			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.expectedError) {
					t.Errorf("CreateAccount() gotError = %v, want %v", err, tt.expectedError)
				}
			}
		})
	}
}
