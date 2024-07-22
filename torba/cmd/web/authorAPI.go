package main

import (
	"errors"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"openSo/internal/postgres"
	"strconv"
)

//author/create GET

func (app *Application) authorForm(w http.ResponseWriter, r *http.Request) error {

	if r.URL.Path != "/author/create" {
		app.clientError(w, http.StatusBadRequest, "Invalid path")
		return errors.New("invalid path")
	}

	err := app.RenderTemplate(w, "../../ui/author/create.html", nil)
	if err != nil {
		app.serverError(w, errors.New("error render html"))
		return err
	}
	return nil
}

//author/create POST

func (app *Application) CreateFullAuthor(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()

	if r.URL.Path != "/author/create" {
		app.clientError(w, http.StatusBadRequest, "Invalid path")
		return errors.New("invalid path")
	}

	author := postgres.CreateFullAuthorParams{
		NickName:       r.FormValue("nickname"),
		Email:          r.FormValue("email"),
		PasswordHash:   app.db.HashPass(r.FormValue("password")),
		Payments:       r.FormValue("payments"),
		Bio:            r.FormValue("bio"),
		Link:           r.FormValue("link"),
		AdditionalInfo: r.FormValue("additional_info"),
	}

	err := app.db.CreateFullAuthorValidator(ctx, &author)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.CreateFullAuthor(ctx, &author)
	if err != nil {
		app.serverError(w, errors.New("failed to create author"))
		return err
	}

	red, err := app.db.GetAuthorByName(ctx, author.NickName)
	if err != nil {
		app.serverError(w, errors.New("failed to get author by nickname"))
	}

	http.Redirect(w, r, "/author/profile/"+strconv.FormatInt(red.AuthorID, 10), http.StatusFound)
	return nil
}

//author/id/profile/delete

func (app *Application) DeleteAuthor(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["author_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid author id")
		return err
	}

	valid := int64(id)

	err = app.db.DeleteAuthor(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("error deleting author"))
		return err
	}

	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

//author/{author_id}/profile

func (app *Application) ProfileAuthor(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["author_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid author id")
		return err
	}
	valid := int64(id)

	a, err := app.db.GetAuthorByID(ctx, valid)
	if err != nil {
		app.clientError(w, http.StatusNotFound, "author not found")
		return err
	}

	templ, err := template.ParseFiles("../../ui/author/profile.gohtml")
	if err != nil {
		app.serverError(w, errors.New("error parsing template"))
		return err
	}
	err = templ.Execute(w, a)
	if err != nil {
		app.serverError(w, errors.New("error executing template"))
		return err
	}

	return nil
}

//author/{author_id}/profile/update

func (app *Application) updateProfileAuthor(w http.ResponseWriter, r *http.Request) error {

	tmpl, err := template.ParseFiles("../../ui/author/update.gohtml")
	if err != nil {
		app.serverError(w, errors.New("error parsing template"))
		return err
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		app.serverError(w, errors.New("error executing template"))
		return err
	}

	return nil
}

func (app *Application) UpdateProfileAuthor(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["author_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid author ID")
		return err
	}

	valid := int64(id)

	author := postgres.UpdateAuthorFullParams{
		AuthorID:       valid,
		NickName:       r.FormValue("nickname"),
		Email:          r.FormValue("email"),
		PasswordHash:   app.db.HashPass(r.FormValue("password")),
		Payments:       r.FormValue("payments"),
		Bio:            r.FormValue("bio"),
		Link:           r.FormValue("link"),
		AdditionalInfo: r.FormValue("additional_info"),
	}

	err = app.db.UpdateAuthorValidator(ctx, &author)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.UpdateAuthor(ctx, &author)
	if err != nil {
		app.serverError(w, errors.New("error updating author"))
	}

	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}
