package web

import "net/http"
import (
	"../api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		header := `<!DOCTYPE html>
		<html>
		<body>`
		body := new(api.Lights).GenerateLightTableHTML()

		footer := `</body>
		</html>`
		html := header + body + footer
		return c.HTML(http.StatusOK, html)
		//https://echo.labstack.com/guide/request
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	//e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":80"))
}
