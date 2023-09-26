package models

import (
	"gorm.io/gorm"
)

type Books struct {
	ID           int64          `gorm:"primaryKey" json:"id"`
	Title        string         `gorm:"varchar(300)" json:"title"`
	Description  string         `gorm:"text" json:"description"`
	BookCategory []BookCategory `gorm:"Foreignkey:BookID;association_foreignkey:ID;" json:"book_category"`
	Keyword      string         `gorm:"text" json:"keyword"`
	Price        float64        `gorm:"type:decimal(14,2)" json:"price"`
	Stock        int64          `gorm:"type:int(5)" json:"stock"`
	Publisher    string         `gorm:"text" json:"publisher"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type BookCategory struct {
	ID         int64 `gorm:"primaryKey" json:"id"`
	BookID     int64 `json:"book_id"`
	CategoryID int64 `json:"category_id"`
}

type Categories struct {
	ID           int64          `gorm:"primaryKey" json:"id"`
	Category     string         `gorm:"text" json:"category"`
	BookCategory []BookCategory `gorm:"Foreignkey:CategoryID;association_foreignkey:ID;" json:"book_category"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
}
