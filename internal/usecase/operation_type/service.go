package operation_type

import (
	"context"
	"github.com/supwr/pismo-transactions/internal/entity"
)

type Service struct {
	repository RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{repository: r}
}

func (o *Service) FindById(ctx context.Context, id int) (*entity.OperationType, error) {
	return o.repository.FindById(ctx, id)
}
