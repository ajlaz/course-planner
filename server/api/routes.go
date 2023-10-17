package api

func (a *API) RegisterRoutes() {
	engine := a.Engine
	engine.PATCH("/courses", a.UpsertCourses)
	engine.GET("/courses/:code", a.GetCourse)
	engine.GET("/courses/hubs/:hubs", a.GetCoursesByHubs)
}
