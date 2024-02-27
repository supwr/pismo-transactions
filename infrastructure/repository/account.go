package repository

import (
	"errors"
	"github.com/supwr/pismo-transactions/entity"
	"gorm.io/gorm"
	"log/slog"
)

type AccountRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewAccountAccountRepository(db *gorm.DB, logger *slog.Logger) *AccountRepository {
	return &AccountRepository{
		db:     db,
		logger: logger,
	}
}

func (r *AccountRepository) Create(account entity.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) FindById(id int) (*entity.Account, error) {
	var account *entity.Account

	if err := r.db.First(&account, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.logger.Error("error finding account", slog.Any("error", err))
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) FindByDocument(document entity.Document) (*entity.Account, error) {
	var account *entity.Account

	if err := r.db.First(&account, "document = ?", document).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		r.logger.Error("error finding account", slog.Any("error", err))
		return nil, err
	}

	return account, nil
}
