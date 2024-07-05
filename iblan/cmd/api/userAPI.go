package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"iblan/cmd/structures"
	"net/http"
	"strconv"
)

//[][][[[[][][][][][][][][][]][][[][][][][[][][]]][][][][][][[][][]]][][]][][][][][][][][]

func (s *APIServer) GetUserByIDAPI(c echo.Context) error {
	id := c.Param("id")
	validID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := s.storage.GetUserByID(validID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (s *APIServer) GetUsersAPI(c echo.Context) error {
	users, err := s.storage.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (s *APIServer) CreateUserHandler(c echo.Context) error {
	//make, create request example of struct, and fill it buy info from request body
	//func new user fill db with our info
	u := new(structures.User)
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user := &structures.User{
		Nickname: u.Nickname,
		Password: u.Password,
		Email:    u.Email,
	}
	if err := s.storage.CreateUser(user.Nickname, user.Password, user.Email); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.Redirect(301, "/")
}

func (s *APIServer) DeleteUserHandler(c echo.Context) error {
	id := c.Param("id")
	validID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := s.storage.DeleteUser(validID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.Redirect(301, "/")
}

func (s *APIServer) UpdateUserHandler(c echo.Context) error {
	// Extract the 'id' parameter from the URL
	id := c.Param("id")
	validID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("invalid id format: %s", id))
	}

	var update structures.User
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("error decoding update data: %v", err))
	}

	user, err := s.storage.UpdateUser(validID, update.Nickname, update.Password, update.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("error updating user: %v", err))
	}

	return c.JSON(http.StatusOK, user)
}

//[][][[[[][][][][][][][][][]][][[][][][][[][][]]][][][][][][[][][]]][][]][][][][][][][][]
