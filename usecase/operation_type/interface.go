package operation_type

import "github.com/supwr/pismo-transactions/entity"

type RepositoryInterface interface {
	FindById(id int) (*entity.OperationType, error)
}
