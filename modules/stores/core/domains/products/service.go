package products

import (
	"context"
	"errors"
)

type Service struct {
	repository Repository
	validator  *Validator
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo, validator: &Validator{}}
}

func (s *Service) CreateProduct(ctx context.Context, product *Product) (ID string, err error) {
	err = s.validator.Validate(product)
	if err != nil {
		return "", err
	}

	ID, err = s.repository.Save(ctx, product)
	if err != nil {
		return "", err
	}

	return ID, nil
}

func (s *Service) GetProductsByStoreID(ctx context.Context, storeID string) (list []*Product, err error) {
	if storeID == "" {
		return nil, errors.New("missing parameter StoreID")
	}

	return s.repository.GetAllByStoreID(ctx, storeID)

}
