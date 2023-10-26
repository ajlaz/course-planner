package postgres

import (
	"github.com/ajlaz/course-planner/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func Register(user *models.User, db *gorm.DB) error {
	inDb := models.User{}

	err := db.First(&inDb, "username = ?", user.Username).Error
	if err == nil {
		return models.ErrUserAlreadyExists
	}

	user.Password = HashPassword(user.Password)
	db.Create(&user)
	return nil

}
