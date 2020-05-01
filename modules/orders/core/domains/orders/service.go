package orders

import "context"

type Service struct {
	repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) CreateOrder(ctx context.Context, order *Order) (ID string, err error) {

	return

}
