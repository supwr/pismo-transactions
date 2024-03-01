package entity

const (
	OperationTypeCashBuy = iota + 1
	OperationTypeInstallmentBuy
	OperationTypeWithdraw
	OperationTypePayment
)

var Operations = map[int]string{
	OperationTypeCashBuy:        "COMPRA A VISTA",
	OperationTypeInstallmentBuy: "COMPRA PARCELADA",
	OperationTypeWithdraw:       "SAQUE",
	OperationTypePayment:        "PAGAMENTO",
}

type OperationType struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
