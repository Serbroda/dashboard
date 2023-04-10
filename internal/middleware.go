package internal

import "github.com/labstack/echo/v4"

func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Response().Header().Set("Cache-Control", "no-cache")
		//c.Response().Header().Set("Expires", time.Now().Add(1*time.Second))

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
