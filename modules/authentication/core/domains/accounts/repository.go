package accounts

import "context"

type Repository interface {
	SaveAccount(ctx context.Context, cre *Account) (ID string, err error)

	GetAccountsByUserName(ctx context.Context, username string) (account *Account, err error)

	GetAccountsByValidationHash(ctx context.Context, hash string) (account *Account, err error)
}

//MuckRepository its a muck depency for testing
type MuckRepository struct {
	SaveAccountFunc           func(ctx context.Context, cre *Account) (ID string, err error)
	GetAccountsByUserNameFunc func(ctx context.Context, username string) (account *Account, err error)
}

func (r *MuckRepository) SaveAccount(ctx context.Context, cre *Account) (ID string, err error) {
	return r.SaveAccountFunc(ctx, cre)
}
func (r *MuckRepository) GetAccountsByUserName(ctx context.Context, username string) (account *Account, err error) {
	return r.GetAccountsByUserNameFunc(ctx, username)
}
