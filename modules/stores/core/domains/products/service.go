package products

import "context"

type Service struct {
	repository Repository
	validator  *Validator
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo}
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
