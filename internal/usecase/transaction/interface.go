//go:generate mockgen -destination=mock/interface.go -source=interface.go -package=mock
package transaction

import (
	"context"
	"github.com/supwr/pismo-transactions/internal/entity"
)

type RepositoryInterface interface {
	Create(ctx context.Context, transaction *entity.Transaction) error
}
