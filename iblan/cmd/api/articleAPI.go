package api

import (
	"bytes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"html/template"
	"iblan/cmd/structures"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var a = godotenv.Load()

var Path = os.Getenv("TEMPLPATH")

func (s *APIServer) articleFull(c echo.Context) error {
	category := strings.Trim(c.Param("category"), " \t\n\r")

	if len(category) > 8 {
		return c.JSON(http.StatusBadRequest, "Invalid category")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	article, err := s.storage.GetArticleFull(category, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Article not found")
	}

	t, err := template.ParseFiles(Path + "/articles/articleUno.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parsing template")
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, article); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error executing template")
	}

	return c.HTML(http.StatusOK, buf.String())
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

func (s *APIServer) articlesByCategory(c echo.Context) error {
	category := strings.Trim(c.Param("category"), " \t\n\r")

	if len(category) > 8 {
		return c.JSON(http.StatusBadRequest, "Invalid category")
	}

	articles, err := s.storage.GetArticlesByCategory(category)
	if err != nil {
		return c.JSON(505, "Error fetching articles")
	}

	if len(articles) == 0 {
		return c.JSON(http.StatusNotFound, "No articles found")
	}
	var t, _ = template.ParseFiles(Path + "/articles/articlesCategory.html")

	data := struct {
		Category string
		Articles []*structures.Article
	}{
		Category: category,
		Articles: articles,
	}

	if err = t.Execute(c.Response().Writer, data); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error executing template")
	}

	return nil
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

func (s *APIServer) createArticleAPI(c echo.Context) error {
	if c.Request().Method != "POST" {
		return c.JSON(http.StatusMethodNotAllowed, "Method not allowed")
	}

	title := strings.Trim(c.FormValue("title"), " \t\n\r")
	category := strings.Trim(c.FormValue("category"), " \t\n\r")
	body := strings.Trim(c.FormValue("body"), " \t\n\r")
	payments := strings.Trim(c.FormValue("payments"), " \t\n\r")
	link := strings.Trim(c.FormValue("link"), " \t\n\r")

	a := &structures.Article{
		Title:    title,
		Category: category,
		Body:     body,
		Payments: payments,
		Link:     link,
	}

	article := structures.NewArticle(a.Title, a.Category, a.Body, a.Payments, a.Link)
	if err := s.storage.CreateArticle(article); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.Redirect(http.StatusFound, "/")
}

func (s *APIServer) createArticleHandler(c echo.Context) error {
	if c.Request().Method != "GET" {
		return c.JSON(http.StatusMethodNotAllowed, "Method not allowed")
	}

	t, err := template.ParseFiles(Path + "/articles/articleCreate.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parsing template")
	}

	err = t.Execute(c.Response(), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error executing template")
	}
	return c.JSON(http.StatusCreated, nil)
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

func (s *APIServer) UpdateArticleAPI(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var a structures.Article
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	error := s.storage.UpdateArticle(id, a.Title, a.Category, a.Body, a.Payments, a.Link)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.Redirect(http.StatusFound, "/")
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

func (s *APIServer) DeleteArticleAPI(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := s.storage.DeleteArticle(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.Redirect(http.StatusFound, "/")
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

func (s *APIServer) GetArticleByIDAPI(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	article, err := s.storage.GetArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, article)
}
