package api

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
)

func (s *APIServer) userFormHandler(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.ErrNotFound
	}
	t, err := template.ParseFiles("/home/voodie/iblan/ui/userForm.html")
	if err != nil {
		return echo.ErrInternalServerError
	}
	t.Execute(c.Response().Writer, nil)
	return nil
}

func (s *APIServer) memberFormHandler(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.ErrNotFound
	}
	t, err := template.ParseFiles("/home/voodie/iblan/ui/memberForm.html")
	if err != nil {
		return echo.ErrInternalServerError
	}
	t.Execute(c.Response().Writer, nil)
	return nil
}
