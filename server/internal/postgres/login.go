package postgres

import (
	"errors"

	"github.com/ajlaz/course-planner/models"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func Login(user *models.User, db *gorm.DB) (*models.User, error) {
	temp := models.User{}
	err := db.First(&temp, "username = ?", user.Username).Error
	if err != nil {
		return nil, err
	}
	if !CheckPasswordHash(user.Password, temp.Password) {
		return nil, errors.New("Incorrect Password")
	}
	return &temp, nil
}
