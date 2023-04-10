package main

import (
	"embed"
	"github.com/Serbroda/dashboard/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	//go:embed frontend/dist
	FrontendDist embed.FS
	//go:embed frontend/dist/index.html
	IndexHTML embed.FS
)

var (
	distDirFS     = echo.MustSubFS(FrontendDist, "frontend/dist")
	distIndexHtml = echo.MustSubFS(IndexHTML, "frontend/dist")
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(internal.NoCache)

	e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)

	e.Static("/static", "static")

	internal.RegisterHandlers(e, internal.Handlers{}, "/api/v1")

	e.Logger.Fatal(e.Start(":8080"))
}
