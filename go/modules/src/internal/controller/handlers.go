package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/vlad-marlo/example/go/modules/src/internal/model"
	"net/http"
)

func (ctrl *Controller) HandlePing(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (ctrl *Controller) HandleGetByID(c echo.Context) error {
	id := c.Param("id")
	for _, item := range ctrl.data {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func (ctrl *Controller) HandleGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.data)
}

func (ctrl *Controller) HandleCreate(c echo.Context) error {
	var req model.PostRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	add := model.StoredData{
		ID:   uuid.NewString(),
		Name: req.Name,
	}
	ctrl.data = append(ctrl.data, add)
	return c.JSON(http.StatusCreated, add)
}
