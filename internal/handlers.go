package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handlers struct {
}

func RegisterHandlers(e *echo.Echo, h Handlers, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/config", h.GetConfig, middlewares...)
}

func (h *Handlers) GetConfig(c echo.Context) error {
	conf, err := ReadConfig("conf.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, conf)
}
