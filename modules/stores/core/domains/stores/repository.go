package stores

import "context"

type Repository interface {
	GetAccountStores(ctx context.Context, store *Store) (stores []Store, err error)

	GetAccountByID(ctx context.Context, ID string) (store *Store, err error)

	Save(ctx context.Context, store *Store) (ID string, err error)

	Update(ctx context.Context, ID string, store *Store) (success bool, err error)
}
