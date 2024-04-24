package repository

import (
	"context"
	"errors"
	"github.com/supwr/pismo-transactions/internal/entity"
	"gorm.io/gorm"
	"log/slog"
)

type AccountRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewAccountRepository(db *gorm.DB, logger *slog.Logger) *AccountRepository {
	return &AccountRepository{
		db:     db,
		logger: logger,
	}
}

func (r *AccountRepository) Create(ctx context.Context, account *entity.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) FindById(ctx context.Context, id int) (*entity.Account, error) {
	var account *entity.Account

	if err := r.db.First(&account, "id = ? and deleted_at is null", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.logger.ErrorContext(ctx, "error finding account", slog.Any("error", err))
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) UpdateAvailableLimit(ctx context.Context, account *entity.Account) error {
	var acc *entity.Account
	return r.db.Model(&acc).Where("id = ?", account.ID).Update("available_credit_limit", account.AvailableCreditLimit).Error
}

func (r *AccountRepository) FindByDocument(ctx context.Context, document entity.Document) (*entity.Account, error) {
	var account *entity.Account

	if err := r.db.First(&account, "document = ? and deleted_at is null", document).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.logger.ErrorContext(ctx, "error finding account", slog.Any("error", err))
		return nil, err
	}

	return account, nil
}
