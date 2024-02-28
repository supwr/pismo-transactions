package repository

import (
	"github.com/supwr/pismo-transactions/entity"
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

func (t *TransactionRepository) Create(transaction *entity.Transaction) error {
	return t.db.Create(transaction).Error
}
