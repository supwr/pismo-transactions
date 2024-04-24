//go:generate mockgen -destination=mock/interface.go -source=interface.go -package=mock
package account

import (
	"context"
	"github.com/supwr/pismo-transactions/internal/entity"
)

type RepositoryInterface interface {
	Create(ctx context.Context, account *entity.Account) error
	UpdateAvailableLimit(ctx context.Context, account *entity.Account) error
	FindById(ctx context.Context, id int) (*entity.Account, error)
	FindByDocument(ctx context.Context, document entity.Document) (*entity.Account, error)
}
