package profile

import "context"

type Store interface {
	StoreProfile(ctx context.Context, profile *Profile) (ID string, err error)

	FindProfileByID(ctx context.Context, ID string) (profile *Profile, err error)
}
