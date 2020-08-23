package profiles

import "context"

type Store interface {
	StoreProfile(ctx context.Context, profile *Profile) (ID string, err error)

	FindProfileByID(ctx context.Context, ID string) (profile *Profile, err error)

	FindProfileByAccountID(ctx context.Context, accountID string) (profile *Profile, err error)

	UpdateProfile(ctx context.Context, ID string, profile *Profile) (success bool, err error)
}
