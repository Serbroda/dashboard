package internal

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Handlers struct {
}

func RegisterHandlers(e *echo.Echo, h Handlers, baseUrl string, middlewares ...echo.MiddlewareFunc) {
	e.GET(baseUrl+"/config", h.GetConfig, middlewares...)
	e.POST(baseUrl+"/config", h.SaveConfig, middlewares...)
	e.GET(baseUrl+"/ping", h.Ping, middlewares...)
}

func (h *Handlers) GetConfig(c echo.Context) error {
	conf, err := ReadConfig[map[string]interface{}]("conf.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, conf)
}

func (h *Handlers) SaveConfig(c echo.Context) error {
	var payload string
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	var conf map[string]interface{}
	err := json.Unmarshal([]byte(payload), &conf)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed unmarshal payload: "+err.Error())
	}
	err = SaveConfig[map[string]interface{}]("conf.json", conf)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to save file: "+err.Error())
	}
	return c.NoContent(http.StatusOK)
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
