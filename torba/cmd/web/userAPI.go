package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"openSo/internal/postgres"
	"strconv"
)

// userForm обробляє GET-запит для створення користувача
func (app *Application) userForm(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := app.Templates(w, r, "userForm", nil)
	if err != nil {
		return err
	}

	return nil
}

// CreateUser обробляє POST-запит для створення користувача
func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	usr := postgres.CreateUserParams{
		Username:     r.FormValue("username"),
		Email:        r.FormValue("email"),
		PasswordHash: app.db.HashPass(r.FormValue("password")),
	}

	err := app.db.CreateUserValidator(ctx, &usr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.CreateUser(ctx, &usr)
	if err != nil {
		app.serverError(w, errors.New("error creating user"))
		return err
	}

	red, err := app.db.GetUserByName(ctx, usr.Username)
	if err != nil {
		app.serverError(w, errors.New("error getting user by username"))
		return err
	}

	http.Redirect(w, r, "/user/profile/"+strconv.FormatInt(red.UserID, 10), http.StatusSeeOther)
	return nil
}

// UserProfile обробляє GET-запит для відображення профілю користувача
func (app *Application) UserProfile(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "incorrect user id")
		return err
	}

	u, err := app.db.GetUserByID(ctx, int64(id))
	if err != nil {
		app.serverError(w, errors.New("error getting user by id"))
		return err
	}

	err = app.Templates(w, r, "userProfile", u)
	if err != nil {
		return err
	}
	return nil
}

// userUpdatePage обробляє GET-запит для відображення сторінки оновлення профілю
func (app *Application) userUpdatePage(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid user id")
		return err
	}

	u, err := app.db.GetUserByID(ctx, int64(id))
	if err != nil {
		app.serverError(w, errors.New("error getting user by id"))
		return err
	}

	err = app.Templates(w, r, "userUpdate", u)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

// UserUpdateProfile обробляє POST-запит для оновлення профілю користувача
func (app *Application) UserUpdateProfile(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid user ID")
		return err
	}

	usr := postgres.UpdateUserParams{
		UserID:       int64(id),
		Username:     r.FormValue("username"),
		Email:        r.FormValue("email"),
		PasswordHash: app.db.HashPass(r.FormValue("password")),
	}

	err = app.db.UpdateUserValidator(ctx, &usr)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.UpdateUser(ctx, &usr)
	if err != nil {
		app.serverError(w, errors.New("error updating user"))
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/user/profile/%d", usr.UserID), http.StatusSeeOther)
	return nil
}

// UserDeleteProfile обробляє POST-запит для видалення профілю користувача
func (app *Application) UserDeleteProfile(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid user id")
		return err
	}

	err = app.db.DeleteUser(ctx, int64(id))
	if err != nil {
		app.serverError(w, errors.New("error deleting user"))
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
