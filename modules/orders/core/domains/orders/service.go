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

	return

}
