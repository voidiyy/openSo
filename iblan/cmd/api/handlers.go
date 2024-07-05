package api

import (
	"github.com/labstack/echo/v4"
	"iblan/cmd/storage"
	"log"
	"net/http"
	"time"
)

//structures

type APIServer struct {
	addr    string
	storage storage.GlobalStorage
}

type APIError struct {
	Error string `json:"error"`
}

func NewAPIServer(addr string, store *storage.PostgresStore) *APIServer {
	return &APIServer{
		addr:    addr,
		storage: store,
	}
}

//operate handlers

func NotFoundHandler(c echo.Context) error {
	return c.NoContent(404)
}

func (s *APIServer) Run() {
	e := echo.New()

	server := &http.Server{
		Addr:         s.addr,
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//how i feel : userForm -> CreateUserHandler

	e.GET("/", s.homeHandler)
	e.GET("/about", s.aboutHandler)

	e.GET("/article/form", s.createArticleHandler)

	e.GET("/article/create", NotFoundHandler)
	e.POST("/article/create", s.createArticleAPI)

	e.GET("/article/:category/:id", s.articleFull)
	e.GET("/article/:category", s.articlesByCategory)

	e.PUT("/article/:id", s.UpdateArticleAPI)
	e.DELETE("/article/:id", s.DeleteArticleAPI)

	//[][][][][][][][][][][][][][[][][][][]]]][]][[][][]]][[][][]][][[]

	//e.GET("/acount/:id", s.accountHandler)
	//e.GET("/acount/:id", s.accountAPI)

	e.GET("/user/form", s.userFormHandler)
	e.POST("/user/create", s.CreateUserHandler)
	e.GET("/user/create", NotFoundHandler)

	e.GET("/user", s.GetUsersAPI)
	e.POST("/user", s.CreateUserHandler)

	e.GET("/user/:id", s.GetUserByIDAPI)
	e.PUT("/user/:id", s.UpdateUserHandler)
	e.DELETE("/user/:id", s.DeleteUserHandler)

	//member[][[][]][][][][][][][][][][][][][]]][][][][][][][][][][][][][]]
	e.GET("/member/form", s.memberFormHandler)
	e.POST("/member/create", s.CreateMemberHandler)
	e.GET("/member/create", NotFoundHandler)

	e.GET("/member", s.GetMembersHandler)

	e.DELETE("/member/:id", s.DeleteMemberHandler)
	e.GET("/member/:id", s.GetMemberByIDHandler)
	e.PUT("/member/:id", s.UpdateMemberHandler)

	//articles system

	//articles create

	//e.GET("/favicon.ico", s.faviconHandler)
	//articles ui

	log.Println("Starting API server at ", s.addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
