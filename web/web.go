package web

import "net/http"
import (
	"../api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strings"
)

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		header := `<!DOCTYPE html>
		<html>
		<body>`
		body := new(api.Lights).GenerateLightTableHTML()
		lightList := new(api.Lights).GetListOfLights()

		var buttonList []string

		for _, v := range lightList {
			buttonTemplate := `<form action="" method="post">
    		<input type="submit" name="nameField" value="nameField" />
			</form>`
			singleButton := strings.Replace(buttonTemplate, "nameField", v.Name(), -1)
			buttonList = append(buttonList, singleButton)
		}

		buttonBody := strings.Join(buttonList, "")
		footer := `</body>
		</html>`
		html := header + body + buttonBody + footer
		return c.HTML(http.StatusOK, html)
		//https://echo.labstack.com/guide/request
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	//e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":80"))
}
