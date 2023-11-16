package api

import (
	"strconv"
	"strings"

	"github.com/ajlaz/course-planner/internal/postgres"
	"github.com/ajlaz/course-planner/models"
	"github.com/gin-gonic/gin"
)

func (a *API) UpsertCourses(c *gin.Context) {
	//check if courses are given as an array if it is, run batch insert, otherwise run single insert

	courses := []models.Course{}
	err := c.ShouldBindJSON(&courses)
	if err != nil {
		course := models.Course{}
		err = c.ShouldBindJSON(&course)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		postgres.SingleInsert(course, a.db)
		c.JSON(200, gin.H{"course": course})
	}
	postgres.BatchInsert(courses, a.db)
	c.JSON(200, gin.H{"courses": courses})

}

func (a *API) GetCourse(c *gin.Context) {
	courseId := c.Param("code")
	course := postgres.Select(courseId, a.db)
	c.JSON(200, gin.H{"course": course})
}

func (a *API) GetCoursesByHubs(c *gin.Context) {
	hubs := c.Param("hubs")
	arr := strings.Split(hubs, ",")
	WRI := strings.Split("Writing, Research, and Inquiry", ",")
	for _, hub := range WRI {
		for i, hub2 := range arr {
			if hub == hub2 {
				arr[i] = "Writing, Research, and Inquiry"
			}
		}
	}

	courses := postgres.SelectByHubs(arr, a.db)
	c.JSON(200, gin.H{"courses": courses})
}

func (a *API) SuggestCourses(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}

	suggestions, hubs := postgres.SuggestCourses(uint(userID), a.db)
	c.JSON(200, gin.H{"suggestions": suggestions, "hubs": hubs})
}
