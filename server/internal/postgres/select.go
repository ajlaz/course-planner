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

	// Build a query to find courses that match any of the hubs
	query := "SELECT * FROM courses WHERE hubs @> ?::text[]"
	db.Raw(query, "{"+strings.Join(hubs, ",")+"}").Scan(&courses)

	// If no courses are found, search for courses that contain at least the first word of each hub
	if len(courses) == 0 {
		for _, hub := range hubs {
			// Extract the first word from the hub
			firstWord := strings.Fields(hub)[0]

			// Use ILIKE to perform a case-insensitive search for courses containing the first word of each hub
			db.Where("lower(hubs::text) LIKE ?", "%"+strings.ToLower(firstWord)+"%").Find(&courses)

			// Break if courses are found
			if len(courses) > 0 {
				break
			}
		}
	}

	return courses
}
