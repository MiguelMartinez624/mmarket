package stores

import "context"

type Repository interface {
	GetStoresByProfileID(ctx context.Context, accountID string) (stores []Store, err error)

	Save(ctx context.Context, store *Store) (ID string, err error)

	Update(ctx context.Context, ID string, store *Store) (success bool, err error)
}