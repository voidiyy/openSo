package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"iblan/cmd/structures"
	"net/http"
	"strconv"
)

/*func (s *APIServer) MemberHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.GetMembersHandler(w, r)
	}
	if r.Method == "POST" {
		return s.CreateMemberHandler(w, r)
	}
	if r.Method == "DELETE" {
		return s.DeleteMemberHandler(w, r)
	}

	return fmt.Errorf("unsupported method: %s", r.Method)
}*/

//[][][[[[][][][][][][][][][]][][[][][][][[][][]]][][][][][][[][][]]][][]][][][][][][][][]

func (s *APIServer) GetMemberByIDHandler(c echo.Context) error {
	var id = c.Param("id")
	validID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid member ID: %s", id)
	}

	member, err := s.storage.GetMemberByID(validID)
	if err != nil {
		return fmt.Errorf("member not found: %s", id)
	}
	return c.JSON(http.StatusOK, member)
}

func (s *APIServer) CreateMemberHandler(c echo.Context) error {
	//make, create request example of struct, and fill it buy info from request body
	//func new member fill db with our info
	u := new(structures.Member)
	if err := c.Bind(&u); err != nil {
		return fmt.Errorf("invalid member request: %s", err)
	}

	member := structures.NewMember(u.Nickname, u.Password, u.Email, u.Category)

	if err := s.storage.CreateMember(member.Nickname, member.Password, member.Email, member.Category); err != nil {
		return fmt.Errorf("cant create member: %s", err)
	}

	return c.Redirect(301, "/")
}

func (s *APIServer) DeleteMemberHandler(c echo.Context) error {
	var id = c.Param("id")
	validID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid member ID: %s", id)
	}
	if err := s.storage.DeleteMember(validID); err != nil {
		return fmt.Errorf("cant delete member: %s", err)
	}
	return c.Redirect(301, "/")
}

func (s *APIServer) GetMembersHandler(c echo.Context) error {
	member, err := s.storage.GetMembers()
	if err != nil {
		return fmt.Errorf("cant get members: %s", err)
	}
	return c.JSON(http.StatusOK, member)
}

func (s *APIServer) UpdateMemberHandler(c echo.Context) error {
	var id = c.Param("id")
	validID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid member ID: %s", id)
	}

	var m structures.Member

	if err := c.Bind(&m); err != nil {
		return fmt.Errorf("invalid member request: %s", err)
	}

	error := s.storage.UpdateMember(validID, m.Nickname, m.Password, m.Email, m.Category)
	if error != nil {
		return fmt.Errorf("cant update member: %s", err)
	}

	return c.Redirect(301, "/")
}
