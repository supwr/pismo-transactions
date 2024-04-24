package account

import (
	"context"
)

type Service struct {
	repository RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{repository: r}
}

func (s *Service) FindById(ctx context.Context, id int) (*Account, error) {
	return s.repository.FindById(ctx, id)
}

func (s *Service) FindByDocument(ctx context.Context, document Document) (*Account, error) {
	return s.repository.FindByDocument(ctx, document)
}

func (s *Service) UpdateCreditLimit(ctx context.Context, account *Account) error {
	return s.repository.UpdateAvailableLimit(ctx, account)
}

func (s *Service) Create(ctx context.Context, account *Account) error {
	exists, err := s.repository.FindByDocument(ctx, account.Document)
	if err != nil {
		return err
	}

	if exists != nil {
		return ErrAccountAlreadyExists
	}

	if err = s.repository.Create(ctx, account); err != nil {
		return err
	}

	return nil
}
