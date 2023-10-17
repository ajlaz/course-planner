package postgres

import (
	"strings"

	"github.com/ajlaz/course-planner/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func Select(code string, db *gorm.DB) models.Course {
	var course models.Course
	db.Where("course_code = ?", code).Find(&course)
	return course
}

func SelectByHubs(hubs []string, db *gorm.DB) []models.Course {
	var courses []models.Course
	query := "SELECT * FROM courses WHERE hubs @> ?::text[]"
	db.Raw(query, "{"+strings.Join(hubs, ",")+"}").Scan(&courses)
	return courses
}
