package api

import (
	"github.com/ajlaz/course-planner/internal/postgres"
	"github.com/ajlaz/course-planner/models"
	"github.com/gin-gonic/gin"
)

func (a *API) AddCoursesToUser(c *gin.Context) {
	req := models.CourseRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	postgres.AddCoursesToUser(req.UserID, req.CourseCodes, a.db)
	c.JSON(200, gin.H{"courses": req.CourseCodes})
}

func (a *API) GetCoursesByUser(c *gin.Context) {
	userID := c.Param("id")
	courses := postgres.GetAllUserCourses(userID, a.db)
	c.JSON(200, gin.H{"courses": courses})
}

func (a *API) Login(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	temp, err := postgres.Login(&user, a.db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": temp})
}

func (a *API) Register(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = postgres.Register(&user, a.db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": user})
}

func (a *API) GetUserRemainingHubs(c *gin.Context) {
	userID := c.Param("id")
	courses := postgres.GetAllUserCourses(userID, a.db)
	hubs := postgres.CalculateRemainingHubs(courses)
	c.JSON(200, gin.H{"hubs": hubs})
}
