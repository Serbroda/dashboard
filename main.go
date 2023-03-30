package main

import (
	"embed"
	"github.com/labstack/echo/v4"
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

	e.StaticFS("/", distDirFS)
	e.FileFS("/", "index.html", distIndexHtml)

	e.Logger.Fatal(e.Start(":8080"))
}
