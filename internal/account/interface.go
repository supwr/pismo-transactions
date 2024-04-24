//go:generate mockgen -destination=mock.go -source=interface.go -package=account
package account

import (
	"context"
)

type RepositoryInterface interface {
	Create(ctx context.Context, account *Account) error
	UpdateAvailableLimit(ctx context.Context, account *Account) error
	FindById(ctx context.Context, id int) (*Account, error)
	FindByDocument(ctx context.Context, document Document) (*Account, error)
}
