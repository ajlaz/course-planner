package api

import (
	"github.com/ajlaz/course-planner/internal/postgres"
	"github.com/ajlaz/course-planner/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	*gin.Engine
	db *gorm.DB
}

func New() *API {
	db := postgres.Connect()
	db.AutoMigrate(&models.Course{}, &models.User{})
	engine := gin.Default()
	engine.Use(CORSMiddleware())
	return &API{
		Engine: engine,
		db:     db,
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.Default()
}
