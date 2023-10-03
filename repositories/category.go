package repositories

import (
	"asis_quest/connection"
	"asis_quest/models"
)

type newCategory struct {
	db connection.DatabaseMain
}

func NewCategory(db connection.DatabaseMain) Category {
	return &newCategory{db: db}
}

func (n *newCategory) Store(insert models.Categories) (*models.Categories, error) {
	var (
		err error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Create(&insert).Error
	if err != nil {
		return nil, err
	}

	return &insert, nil
}

func (n *newCategory) StoreBookCategory(insert models.BookCategory) (*models.BookCategory, error) {
	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	if err := database.Create(&insert).Error; err != nil {
		return nil, err
	}

	return &insert, nil
}

func (n *newCategory) FindOne(where models.Categories) (*models.Categories, error) {
	var (
		category *models.Categories
		err      error
	)
	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("category = ?", where.Category).First(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (n *newCategory) FindBookCategory(where models.BookCategory) ([]models.BookCategory, error) {
	var (
		bookCategory []models.BookCategory
		err          error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("book_id = ?", where.BookID).Find(&bookCategory).Error
	if err != nil {
		return nil, err
	}

	return bookCategory, nil
}

func (n *newCategory) FindID(where models.Categories) (*models.Categories, error) {
	var (
		category *models.Categories
		err      error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("id = ?", where.ID).Find(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}
