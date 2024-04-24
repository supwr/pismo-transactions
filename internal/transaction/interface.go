//go:generate mockgen -destination=mock.go -source=interface.go -package=transaction
package transaction

import (
	"context"
)

type RepositoryInterface interface {
	Create(ctx context.Context, transaction *Transaction) error
}
