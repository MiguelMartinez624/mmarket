package stores

import "context"

type Repository interface {
	GetStoresByProfileID(ctx context.Context, profileID string) (list []Store, err error)
	GetStoreByIDAndProfileID(ctx context.Context, storeID string, profileID string) (item *Store, err error)

	Save(ctx context.Context, store *Store) (ID string, err error)

	Update(ctx context.Context, ID string, store *Store) (success bool, err error)

	GetByID(ctx context.Context, ID string) (product *Store, err error)
}
