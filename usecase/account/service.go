package account

import (
	"errors"
	"github.com/supwr/pismo-transactions/entity"
)

type Service struct {
	repository RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{repository: r}
}

func (s *Service) FindById(id int) (*entity.Account, error) {
	return s.repository.FindById(id)
}

func (s *Service) Create(account entity.Account) error {
	exists, err := s.repository.FindByDocument(account.Document)
	if err != nil {
		return err
	}

	if exists != nil {
		return errors.New("There's already an account with this document")
	}

	if err = s.repository.Create(account); err != nil {
		return err
	}

	return nil
}
