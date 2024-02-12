package controller

import (
	"github.com/labstack/echo"
	"github.com/vlad-marlo/example/go/six/internal/model"
	"net/http"
)

func (ctrl *Controller) HandlePing(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (ctrl *Controller) HandleGetByID(c echo.Context) error {
	resp, err := ctrl.srv.HandleGetByID(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) HandleGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.srv.HandleGetAll())
}

func (ctrl *Controller) HandleCreate(c echo.Context) error {
	var req model.PostRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	resp := ctrl.srv.HandleCreate(req.Name)
	return c.JSON(http.StatusCreated, resp)
}
