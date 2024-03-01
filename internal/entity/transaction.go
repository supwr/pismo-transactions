package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID              int             `json:"id" gorm:"primaryKey"`
	AccountID       int             `json:"account_id"`
	OperationTypeID int             `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
	OperationDate   time.Time       `json:"operation_date"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       *time.Time      `json:"updated_at"`
	DeletedAt       *time.Time      `json:"deleted_at"`
}
