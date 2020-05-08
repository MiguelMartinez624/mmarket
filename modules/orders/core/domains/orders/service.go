package orders

import (
	"context"
	"errors"
)

type Service struct {
	repository Repository
	validator  Validator
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo, validator: Validator{}}
}

func (s *Service) CreateOrder(ctx context.Context, order *Order) (created *Order, err error) {

	err = s.validator.Validate(order)
	if err != nil {
		return nil, err
	}

	// set to pending
	order.Status = Pending

	ID, err := s.repository.SaveOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// Get the recent created order
	created, err = s.repository.GetOrderByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return created, nil

}

func (s *Service) GetOrderByStoreID(ctx context.Context, storeID string) (list []Order, err error) {
	if storeID == "" {
		return nil, errors.New("Missing storeID")
	}

	list, err = s.repository.GetOrdersByStoreID(ctx, storeID)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Service) GetOrderByCostumerID(ctx context.Context, costumerID string) (list []Order, err error) {
	if costumerID == "" {
		return nil, errors.New("Missing costumerID")
	}

	list, err = s.repository.GetOrdersByCostumerID(ctx, costumerID)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Service) UpdateOrder(ctx context.Context, orderID string, order Order) (ok bool, err error) {
	return
}
