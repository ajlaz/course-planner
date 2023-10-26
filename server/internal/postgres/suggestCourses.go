package postgres

import (
	"strings"

	"github.com/ajlaz/course-planner/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func GetAllUserCourses(userId string, db *gorm.DB) []models.Course {
	user := models.User{}
	db.First(&user, userId)
	temp := strings.Split(user.Courses, ",")
	courses := []models.Course{}
	for _, coursename := range temp {
		course := models.Course{}
		db.First(&course, "course_code = ?", coursename)
		courses = append(courses, course)
	}
	return courses
}

func CalculateRemainingHubs(courses []models.Course) map[string]int {
	hubs := models.GetHubs()
	for _, course := range courses {
		for _, hub := range course.Hubs {
			hubs[hub]--
		}
	}
	return hubs
}

func AddCoursesToUser(userID uint, courses string, db *gorm.DB) {
	user := models.User{}
	db.First(&user, userID)
	user.Courses += courses
	temp := strings.Split(user.Courses, ",")
	//remove any duplicates from temp
	temp = RemoveDuplicates(temp)
	temp = EnsureExists(temp, db)
	user.Courses = strings.Join(temp, ",") + ","
	//remove all spaces after commas
	user.Courses = strings.ReplaceAll(user.Courses, ", ", ",")
	user.Courses = strings.ReplaceAll(user.Courses, ",,", ",")

	db.Save(&user)

}

func RemoveDuplicates(courses []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range courses {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func EnsureExists(courses []string, db *gorm.DB) []string {
	var temp []string
	for _, course := range courses {
		c := models.Course{}
		db.First(&c, "course_code = ?", course)
		if c.CourseCode != "" {
			temp = append(temp, course)
		}
	}
	return temp

}
