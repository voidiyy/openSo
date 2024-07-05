package api

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
)

func (s *APIServer) homeHandler(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.ErrNotFound
	}
	t, err := template.ParseGlob("/home/voodie/iblan/ui/index.html")
	if err != nil {
		return echo.ErrInternalServerError
	}
	t.Execute(c.Response().Writer, nil)
	return nil

}

func (s *APIServer) aboutHandler(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.ErrNotFound
	}
	t, err := template.ParseFiles("/home/voodie/iblan/ui/about.html")
	if err != nil {
		return echo.ErrInternalServerError
	}
	t.Execute(c.Response().Writer, nil)
	return nil

}
