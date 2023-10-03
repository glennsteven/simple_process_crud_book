package repositories

import (
	"asis_quest/connection"
	"asis_quest/models"
)

type newBook struct {
	db connection.DatabaseMain
}

func NewBook(db connection.DatabaseMain) Book {
	return &newBook{db: db}
}

func (n *newBook) Find(where []models.Books) ([]models.Books, error) {
	var (
		err error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Find(&where).Error
	if err != nil {
		return nil, err
	}

	return where, nil
}

func (n *newBook) FindOne(where models.Books) (*models.Books, error) {
	var (
		result *models.Books
		err    error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("title = ?", where.Title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (n *newBook) Store(insert models.Books) (*models.Books, error) {
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

func (n *newBook) FindID(where models.Books) (*models.Books, error) {
	var (
		result *models.Books
		err    error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("id = ?", where.ID).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (n *newBook) Update(id int64, update models.Books) (*models.Books, error) {
	var (
		result *models.Books
		err    error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("id = ?", id).Updates(&update).Error
	if err != nil {
		return nil, err
	}

	result = &update

	return result, nil
}
