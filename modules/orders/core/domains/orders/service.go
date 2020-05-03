package orders

import "context"

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
	return
}
