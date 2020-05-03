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

func (s *Service) UpdateProduct(ctx context.Context, ID string, product *Product) (success bool, err error) {
	if ID == "" {
		return false, errors.New("missing parameter StoreID")
	}
	if product == nil {
		return false, errors.New("missing parameter product")
	}
	err = s.validator.Validate(product)
	if err != nil {
		return false, err
	}

	_, err = s.repository.GetByID(ctx, ID)
	if err != nil {
		return false, err
	}

	success, err = s.repository.Update(ctx, ID, product)
	if err != nil {
		return false, err
	}

	return
}

func (s *Service) CheckAvailability(ctx context.Context, productID string, quantity int) (ok bool, err error) {

	p, err := s.repository.GetByID(ctx, productID)
	if err != nil {
		return false, err
	}

	if p.Stock < quantity {
		return false, err
	}

	return true, nil
}
