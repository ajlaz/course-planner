package api

func (a *API) RegisterRoutes() {
	engine := a.Engine
	engine.PATCH("/courses", a.UpsertCourses)
	engine.GET("/courses/:code", a.GetCourse)
	engine.GET("/courses/hubs/:hubs", a.GetCoursesByHubs)
	engine.POST("/users/courses", a.AddCoursesToUser)
	engine.GET("/users/:id/courses", a.GetCoursesByUser)
	engine.POST("/register", a.Register)
	engine.POST("/login", a.Login)
	engine.GET("/users/:id/hubs", a.GetUserRemainingHubs)
}
