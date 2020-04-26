package stores

import (
	"context"
	"errors"
	"fmt"
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

func (s *Service) GetStoreByID(ctx context.Context, ID string) (store *Store, err error) {
	if ID == "" {
		return nil, errors.New("missing parameter StoreID")
	}

	store, err = s.repo.GetByID(ctx, ID)
	fmt.Println(store)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetStoreByIDAndProfileID(ctx context.Context, storeID string, profileID string) (store *Store, err error) {
	if storeID == "" {
		return nil, ErrMissingParam{Param: "storeID"}
	}
	if profileID == "" {
		return nil, ErrMissingParam{Param: "profileID"}
	}

	store, err = s.repo.GetStoreByIDAndProfileID(ctx, storeID, profileID)
	return
}
