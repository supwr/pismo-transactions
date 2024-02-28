package transaction

import "github.com/supwr/pismo-transactions/entity"

type RepositoryInterface interface {
	Create(transaction *entity.Transaction) error
}
