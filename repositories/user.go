package repositories

import (
	"asis_quest/connection"
	"asis_quest/models"
	"errors"
	"gorm.io/gorm"
)

type newUser struct {
	db connection.DatabaseMain
}

func NewUser(db connection.DatabaseMain) User {
	return &newUser{db: db}
}

func (n *newUser) FindUser(where models.User) (*models.User, error) {
	var (
		result *models.User
		err    error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = database.Where("user_name = ?", where.UserName).First(&result).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			msg := errors.New("not found")
			return nil, msg
		default:
			return nil, err
		}
	}

	return result, nil
}

func (n *newUser) InsertUser(insert models.User) error {
	var (
		err error
	)

	database, err := n.db.ConnectDatabase()
	if err != nil {
		return err
	}

	err = database.Create(&insert).Error
	if err != nil {
		return err
	}

	return nil
}
