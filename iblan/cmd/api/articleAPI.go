package api

import (
	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"iblan/cmd/structures"
	contentarticle "iblan/ui/contentArticle"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var a = godotenv.Load()

var Path = os.Getenv("TEMPLPATH")

// /slip/:category/:id
func (s *APIServer) articleFull(c echo.Context) error {
	category := strings.Trim(c.Param("category"), " \t\n\r")

	if len(category) > 20 {
		return c.JSON(http.StatusBadRequest, "Invalid category")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	article, err := s.storage.GetArticleFull(category, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return Renderer(c, contentarticle.SingleArticle(article))

}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

// /slip/:category
func (s *APIServer) articlesByCategory(c echo.Context) error {
	category := strings.Trim(c.Param("category"), " \t\n\r")

	if len(category) > 20 {
		return c.JSON(http.StatusBadRequest, "Invalid category")
	}

	articles, err := s.storage.GetArticlesByCategory(category)
	if err != nil {
		return c.JSON(505, "Error fetching articles")
	}

	return Renderer(c, contentarticle.ManyArticles(articles))
}

func Renderer(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

// /e.POST("/slip/create
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

	// !!!!!!!!! add templ form
	article := structures.NewArticle(a.Title, a.Category, a.Body, a.Payments, a.Link)
	if err := s.storage.CreateArticle(article); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return Renderer(c, contentarticle.SingleArticle(article))
}

// e.GET("/slip/form
func (s *APIServer) createArticleHandler(c echo.Context) error {
	if c.Request().Method != "GET" {
		return c.JSON(http.StatusMethodNotAllowed, "Method not allowed")
	}

	return Renderer(c, contentarticle.FormArticles())
}

//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

func (s *APIServer) UpdateArticleAPI(c echo.Context) error {
	title := strings.Trim(c.Param("title"), " \t\n\r")

	var a structures.Article
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updated, error := s.storage.UpdateArticle(title, a.Title, a.Category, a.Body, a.Payments, a.Link)

	if error != nil {
		return c.JSON(http.StatusInternalServerError, error)
	}
	newArticle, err := s.storage.GetArticleByID(updated.ID)
	if err != nil {
		return err
	}
	return Renderer(c, contentarticle.SingleArticle(newArticle))
}

func (s *APIServer) UpdateArticleHandler(c echo.Context) error {
	if c.Request().Method != "GET" {
		return c.JSON(http.StatusMethodNotAllowed, "Method not allowed")
	}
	return Renderer(c, contentarticle.FormArticles())
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
	article, err := s.storage.GetArticleByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return Renderer(c, contentarticle.SingleArticle(article))
}

//revrite all in templ last is this |^
