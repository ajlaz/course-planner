package postgres

import (
	"strings"

	"github.com/ajlaz/course-planner/models"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func GetAllUserCourses(userId string, db *gorm.DB) string {
	//get courses for a user with id userId as a string array only

	var temp string
	db.Raw("SELECT courses FROM users WHERE id = ?", userId).Scan(&temp)
	return temp
	// var courses []models.Course
	// for _, coursename := range temp {
	// 	course := Select(coursename, db)
	// 	courses = append(courses, course)
	// }
	// return courses
}

func calculateRemainingHubs(courses []models.Course) map[string]int {
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
	user.Courses = strings.Join(temp, ",") + ","

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
