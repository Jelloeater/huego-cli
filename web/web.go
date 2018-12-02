package web

import "net/http"
import (
	"../api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
	"strconv"
	"strings"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["buttoncall"] = GenerateButtons()
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func GenerateButtons() string {
	lightList := new(api.Lights).GetListOfLights()
	var buttonList []string

	for _, v := range lightList {
		buttonTemplate := `<form action="/light/on/id" method="post">
    		<input type="submit" name="nameField" value="nameField" />
			</form>`
		singleButton := strings.Replace(buttonTemplate, "nameField", v.Name(), 1)
		buttonList = append(buttonList, singleButton)
	}

	return strings.Join(buttonList, "")
}

func StartServer() {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		l := api.Lights{}
		lOut := l.GenerateLightTableHTML()

		return c.Render(http.StatusOK, "main.html", map[string]interface{}{
			"lights":  template.HTML(lOut),
			"buttons": template.HTML(GenerateButtons()),
		})
	})

	e.POST("/light/on/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logrus.Error("Bad ID")
		} else {
			l := api.Light{}
			l = l.GetLight(id)
			l.TurnOn()
		}
		return c.HTML(http.StatusOK, strconv.Itoa(id))
	})

	e.POST("/light/off/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logrus.Error("Bad ID")
		} else {
			l := api.Light{}
			l = l.GetLight(id)
			l.TurnOff()
		}
		return c.HTML(http.StatusOK, strconv.Itoa(id))
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	//e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":80"))
}
