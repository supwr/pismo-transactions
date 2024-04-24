package transaction

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
)

type Repository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewRepository(db *gorm.DB, logger *slog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (t *Repository) Create(ctx context.Context, transaction *Transaction) error {
	return t.db.Create(transaction).Error
}
