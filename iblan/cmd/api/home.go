package api

import (
	"github.com/labstack/echo/v4"
	"iblan/ui/contentIndex"
	"net/http"
)

func (s *APIServer) homeHandler(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.ErrNotFound
	}

	return Renderer(c, contentIndex.HomeIndex())
}

func (s *APIServer) aboutHandler(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.ErrNotFound
	}

	return Renderer(c, contentIndex.AboutIndex())

}
