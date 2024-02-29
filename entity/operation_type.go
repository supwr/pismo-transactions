package entity

const (
	OperationTypeCashBuy = iota + 1
	OperationTypeInstallmentBuy
	OperationTypeWithdraw
	OperationTypePayment
)

type OperationType struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
