package routes

import (
	"github.com/labstack/echo"
	"plm/controller"
)

type Routing struct {
	routes controller.Controller
}

type RoutingInterface interface {
	GetRoutes() *echo.Echo
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()
	
	//user
	e.POST("/tasks/user/register/", Routing.routes.PostsRegis)
	e.POST("/tasks/user/login/", Routing.routes.GetLogin)
	e.PUT("/tasks/user/edit/", Routing.routes.UpEdit)

	//project
	e.POST("/tasks/project/create/", Routing.routes.PostsCreate)
	e.POST("/tasks/project/view/", Routing.routes.ViewProject)
	e.PUT("/tasks/project/edit/", Routing.routes.EditProject)

	return e
}




