package entity

type OperationType struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
