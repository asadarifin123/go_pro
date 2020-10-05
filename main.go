package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,

) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

//func (io.Writer, string, interface{}, echo.Context) error

func main() {

	e := echo.New()

	e.Static("/static", "assets")

	e.Renderer = NewRenderer("views/*.html", true)

	e.GET("/index", func(c echo.Context) error {
		data := M{"message": "HOME"}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.GET("/login", func(c echo.Context) error {
		data := M{"message": "LOGIN"}
		return c.Render(http.StatusOK, "login.html", data)
	})

	e.GET("/postmortem", func(c echo.Context) error {
		data := M{"message": "POSTMORTEM"}
		return c.Render(http.StatusOK, "postmortem.html", data)
	})

	e.GET("/dashboard", func(c echo.Context) error {
		data := M{"message": "DASHBOARD"}
		return c.Render(http.StatusOK, "dashboard.html", data)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
