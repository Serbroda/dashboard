package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct {
}

type Data struct {
	Title string `json:"title"`
}

func RegisterHandlers(e *echo.Echo, h Handlers, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/config", h.GetConfig, middlewares...)
}

func (h *Handlers) GetConfig(c echo.Context) error {
	return c.JSON(http.StatusOK, Data{
		Title: "Test",
	})
}
