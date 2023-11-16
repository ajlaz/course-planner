package main

import "github.com/ajlaz/course-planner/api"

func main() {
	a := api.New()
	a.RegisterRoutes()
	a.Run()

}
