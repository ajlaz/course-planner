package postgres

import (
	"github.com/ajlaz/course-planner/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func SingleInsert(course models.Course, db *gorm.DB) {
	db.Create(&course)
}

func BatchInsert(courses []models.Course, db *gorm.DB) {
	db.CreateInBatches(&courses, 100)
}
