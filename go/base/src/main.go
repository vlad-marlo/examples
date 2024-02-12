package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

type PostRequest struct {
	Name string `json:"name"`
}

type StoredData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	var data = make([]StoredData, 0, 10)

	e.POST("/", func(c echo.Context) error {
		var req PostRequest
		if err := c.Bind(&req); err != nil {
			return err
		}
		add := StoredData{
			ID:   uuid.NewString(),
			Name: req.Name,
		}
		data = append(data, add)
		return c.JSON(http.StatusCreated, add)
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		for _, item := range data {
			if item.ID == id {
				return c.JSON(http.StatusOK, item)
			}
		}
		return c.NoContent(http.StatusNotFound)
	})
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, data)
	})

	if err := e.Start("localhost:8080"); err != nil {
		panic(err)
	}
}
