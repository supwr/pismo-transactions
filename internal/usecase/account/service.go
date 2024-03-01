package account

import (
	"github.com/supwr/pismo-transactions/internal/entity"
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

func (s *Service) FindByDocument(document entity.Document) (*entity.Account, error) {
	return s.repository.FindByDocument(document)
}

func (s *Service) Create(account *entity.Account) error {
	exists, err := s.repository.FindByDocument(account.Document)
	if err != nil {
		return err
	}

	if exists != nil {
		return ErrAccountAlreadyExists
	}

	if err = s.repository.Create(account); err != nil {
		return err
	}

	return nil
}
