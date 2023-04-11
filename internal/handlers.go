package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Handlers struct {
}

func RegisterHandlers(e *echo.Echo, h Handlers, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/config", h.GetConfig, middlewares...)
	e.GET(baseUrl+"/ping", h.Ping, middlewares...)
}

func (h *Handlers) GetConfig(c echo.Context) error {
	conf, err := ReadConfig[map[string]interface{}]("conf.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, conf)
}

func (h *Handlers) Ping(c echo.Context) error {
	target := c.QueryParam("target")
	return c.NoContent(pingURL(target))
}

func pingURL(url string) int {
	url = strings.TrimSpace(url)
	res, err := http.Get(url)
	if err != nil {
		return http.StatusNotFound
	}
	return res.StatusCode
}
