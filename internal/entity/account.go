package entity

import "time"

type Document string

type Account struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Document  Document   `json:"document"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
