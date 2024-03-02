package repository

import (
	"context"
	"github.com/supwr/pismo-transactions/internal/entity"
	"gorm.io/gorm"
	"log/slog"
)

type TransactionRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewTransactionRepository(db *gorm.DB, logger *slog.Logger) *TransactionRepository {
	return &TransactionRepository{
		db:     db,
		logger: logger,
	}
}

func (t *TransactionRepository) Create(ctx context.Context, transaction *entity.Transaction) error {
	return t.db.Create(transaction).Error
}
