package models

import "time"

type Books struct {
	Id          int64      `gorm:"primaryKey" json:"id"`
	CategoryID  int64      `gorm:"type:int" json:"category_id"`
	Title       string     `gorm:"varchar(300)" json:"title"`
	Description string     `gorm:"text" json:"description"`
	Category    string     `gorm:"text" json:"category"`
	Keyword     string     `gorm:"text" json:"keyword"`
	Price       float64    `gorm:"type:decimal(14,2)" json:"price"`
	Stock       int64      `gorm:"type:int(5)" json:"stock"`
	Publisher   string     `gorm:"text" json:"publisher"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Categories struct {
	Id        int64      `gorm:"primaryKey" json:"id"`
	Category  string     `gorm:"text" json:"category"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
