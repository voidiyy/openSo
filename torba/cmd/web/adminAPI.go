package main

import (
	"errors"
	"net/http"
)

func (app *Application) panel(w http.ResponseWriter, r *http.Request) error {
	err := app.Templates(w, r, "panel", nil)
	if err != nil {
		app.serverError(w, errors.New("error render page"))
	}
	return err
}

func (app *Application) listUsers(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	user, err := app.db.ListUserName(ctx)
	if err != nil {
		app.serverError(w, errors.New("error list users by name"))
		return err
	}

	err = app.Templates(w, r, "panelUsers", user)
	if err != nil {
		app.serverError(w, errors.New("error listing users by name"))
		return err
	}

	return nil
}

func (app *Application) listAuthors(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	authors, err := app.db.ListAuthorsByNick(ctx)
	if err != nil {
		app.serverError(w, errors.New("error list authors by nick"))
		return err
	}

	err = app.Templates(w, r, "panelAuthors", authors)
	if err != nil {
		app.serverError(w, errors.New("error listing authors by name"))
		return err
	}
	return nil
}
