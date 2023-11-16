package postgres

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

func getAllCourses(db *gorm.DB) []models.Course {
	courses := []models.Course{}
	db.Find(&courses)
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

func DeleteCourseFromUser(userid uint, courseCode string, db *gorm.DB) {
	user := models.User{}
	db.First(&user, userid)
	temp := strings.Split(user.Courses, ",")
	for i, course := range temp {
		if course == courseCode {
			temp = append(temp[:i], temp[i+1:]...)
		}
	}
	user.Courses = strings.Join(temp, ",") + ","
	user.Courses = strings.ReplaceAll(user.Courses, ", ", ",")
	user.Courses = strings.ReplaceAll(user.Courses, ",,", ",")
	db.Save(&user)
}

func SuggestCourses(userID uint, db *gorm.DB) ([]models.Course, map[string]int) {
	userCourses := GetAllUserCourses(strconv.Itoa(int(userID)), db)
	hubs := CalculateRemainingHubs(userCourses)
	courses := userCourses
	i := 0

	//pick a random hub from the remaining hubs and get all courses that satisfy that hub
	for !hubsSatisfied(hubs) {
		fmt.Println("iterations: " + strconv.Itoa(i))
		i += 1
		remainingHubs := []string{}
		for hub, value := range hubs {
			if value > 0 {
				remainingHubs = append(remainingHubs, hub)
			}
		}
		randomHubs := randomHubsSubset(remainingHubs)
		selectedCourses := SelectByHubs(randomHubs, db)
		for _, course := range selectedCourses {
			if canAddCourse(course, courses) {
				fmt.Println(course.CourseCode)
				courses = append(courses, course)
				break
			}
			hubs = CalculateRemainingHubs(courses)
			printHubStatus(hubs)

		}
		//if I is greater than 100, iterate through each remaining hub and add a course that satisfies it
		if i > 100 {
			for _, hub := range remainingHubs {
				fmt.Println("checking: " + hub)
				selectedCourses := SelectByHubs([]string{hub}, db)
				for _, course := range selectedCourses {
					if canAddCourse(course, courses) {
						fmt.Println(course.CourseCode)
						courses = append(courses, course)
						break
					}
					hubs = CalculateRemainingHubs(courses)
					printHubStatus(hubs)

				}
			}
		}

	}
	//set courses to all courses except user courses
	courses = courses[len(userCourses):]
	//remove any hub from hubs that contains "Pathway"
	for hub := range hubs {
		if strings.Contains(hub, "Pathway") {
			delete(hubs, hub)
		}
	}
	return courses, hubs
}

func randomHubsSubset(hubs []string) []string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3) + 1 // select 1 to 3 hubs randomly
	if n > len(hubs) {
		n = len(hubs)
	}
	shuffledHubs := shuffleHubs(hubs)
	return shuffledHubs[:n]
}

func shuffleHubs(hubs []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hubs), func(i, j int) {
		hubs[i], hubs[j] = hubs[j], hubs[i]
	})
	return hubs
}

// Function to check if all hubs are satisfied
func hubsSatisfied(hubs map[string]int) bool {
	for _, v := range hubs {
		if v > 0 {
			return false
		}
	}
	return true
}

// Function to print courses
func printCourses(courses []models.Course) {
	for _, course := range courses {
		println(course.CourseCode)
	}
}

// Function to print hub status
func printHubStatus(hubs map[string]int) {
	for hub, value := range hubs {
		println(hub, value)
	}
}

// Function to check if a course can be added
func canAddCourse(course models.Course, courses []models.Course) bool {
	for _, c := range courses {
		if c.CourseCode == course.CourseCode {
			return false
		}
	}
	return true
}
