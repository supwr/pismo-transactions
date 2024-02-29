//go:generate mockgen -destination=mock/interface.go -source=interface.go -package=mock
package transaction

import "github.com/supwr/pismo-transactions/entity"

type RepositoryInterface interface {
	Create(transaction *entity.Transaction) error
}
