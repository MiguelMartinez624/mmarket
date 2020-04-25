package stores

import (
	"context"
	"errors"
)

type Service struct {
	repo      Repository
	validator *Validator
}

func NewService(storeRepo Repository) *Service {
	return &Service{repo: storeRepo, validator: &Validator{}}
}

func (s *Service) CreateStore(ctx context.Context, store *Store) (ID string, err error) {
	err = s.validator.Validate(store)
	if err != nil {
		return "", err
	}

	ID, err = s.repo.Save(ctx, store)
	if err != nil {
		return "", err
	}

	return ID, nil
}

func (s *Service) GetUserStores(ctx context.Context, profileID string) (list []Store, err error) {
	if profileID == "" {
		return nil, errors.New("no profile")
	}

	list, err = s.repo.GetStoresByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	//If list its null return a empty for beer ux use
	if list == nil {
		return []Store{}, nil
	}

	return
}
