package externals

import (
	"context"

	"github.com/miguelmartinez624/mmarket/modules/stores/core/dto"
)

type ProfileModule interface {
	GetProfileByID(ctx context.Context, ID string) (profile *dto.Profile, err error)
}
