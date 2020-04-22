package stores

import "context"

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
