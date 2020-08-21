package accounts

import "context"

type MockRepository struct {
	SaveHandlerFunc   func() (ID string, err error)
	UpdateHandlerFunc func() (success bool, err error)
	GetByUserNameFunc func() (account *Account, err error)
	GetByHashFunc     func() (account *Account, err error)
}

func (m MockRepository) SaveAccount(ctx context.Context, cre *Account) (ID string, err error) {
	return m.SaveHandlerFunc()
}

func (m MockRepository) UpdateAccount(ctx context.Context, ID string, cre *Account) (success bool, err error) {
	return m.UpdateHandlerFunc()
}

func (m MockRepository) GetAccountsByUserName(ctx context.Context, username string) (account *Account, err error) {
	return m.GetByUserNameFunc()
}

func (m MockRepository) GetAccountsByValidationHash(ctx context.Context, hash string) (account *Account, err error) {
	return m.GetByHashFunc()
}
