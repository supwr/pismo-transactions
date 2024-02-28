package operation_type

import "github.com/supwr/pismo-transactions/entity"

type Service struct {
	repository RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{repository: r}
}

func (o *Service) FindById(id int) (*entity.OperationType, error) {
	return o.repository.FindById(id)
}
