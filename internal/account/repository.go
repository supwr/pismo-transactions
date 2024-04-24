package account

import (
	"context"
	"errors"
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

func (r *Repository) Create(ctx context.Context, account *Account) error {
	return r.db.Create(account).Error
}

func (r *Repository) FindById(ctx context.Context, id int) (*Account, error) {
	var account *Account

	if err := r.db.First(&account, "id = ? and deleted_at is null", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.logger.ErrorContext(ctx, "error finding account", slog.Any("error", err))
		return nil, err
	}

	return account, nil
}

func (r *Repository) UpdateAvailableLimit(ctx context.Context, account *Account) error {
	var acc *Account
	return r.db.Model(&acc).Where("id = ?", account.ID).Update("available_credit_limit", account.AvailableCreditLimit).Error
}

func (r *Repository) FindByDocument(ctx context.Context, document Document) (*Account, error) {
	var account *Account

	if err := r.db.First(&account, "document = ? and deleted_at is null", document).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.logger.ErrorContext(ctx, "error finding account", slog.Any("error", err))
		return nil, err
	}

	return account, nil
}
