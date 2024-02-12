package controller

import (
	"github.com/labstack/echo"
	"github.com/vlad-marlo/example/go/six/internal/service"
)

type Controller struct {
	echo *echo.Echo
	srv  *service.Service
}

func New() *Controller {
	c := &Controller{
		echo: echo.New(),
		srv:  service.New(),
	}
	c.configureRoutes()
	return c
}

func (ctrl *Controller) configureRoutes() {
	ctrl.echo.POST("/", ctrl.HandleCreate)
	ctrl.echo.GET("/ping", ctrl.HandlePing)
	ctrl.echo.GET("/:id", ctrl.HandleGetByID)
	ctrl.echo.GET("/", ctrl.HandleGetAll)
}

func (ctrl *Controller) Start() error {
	return ctrl.echo.Start("localhost:8080")
}
