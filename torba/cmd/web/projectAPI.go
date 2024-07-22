package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"openSo/internal/postgres"
	"strconv"
)

///categories/project/{project_id}

func (app *Application) projectByID(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["project_id"])

	if err != nil {
		app.clientError(w, http.StatusBadRequest, "invalid project id")
		return err
	}
	valid := int32(id)

	prj, err := app.db.GetProjectByID(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("project not found"))
		return err
	}

	err = app.RenderTemplate(w, "", prj)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

///categories/project/create GET

func (app *Application) CreateProjectForm(w http.ResponseWriter, r *http.Request) error {
	err := app.RenderTemplate(w, "", nil)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}
	return nil
}

///categories/project/create POST

func (app *Application) CreateProject(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	prj := postgres.CreateProjectParams{
		AuthorID:    1,
		Title:       r.FormValue("title"),
		Category:    r.FormValue("category"),
		Subcategory: r.FormValue("subcategory"),
		Description: r.FormValue("description"),
		Link:        r.FormValue("link"),
		Details:     r.FormValue("details"),
		Payments:    r.FormValue("payments"),
		FundingGoal: r.FormValue("funding_goal"),
	}

	err := app.db.CreateProjectValidator(ctx, &prj)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.CreateProject(ctx, &prj)
	if err != nil {
		app.serverError(w, errors.New("project creation failed"))
		return err
	}

	id, err := app.db.GetProjectByTitle(ctx, prj.Title)
	if err != nil {
		app.serverError(w, errors.New("project not found"))
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/categories/project/%d", id.ProjectID), http.StatusFound)
	return nil
}

func (app *Application) ListProjectsByCategory(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()

	vars := mux.Vars(r)
	category := vars["category_name"]

	err := app.db.CategoryExistsPrj(ctx, category)
	if err != nil {
		app.serverError(w, errors.New("category not found"))
		return err
	}

	projects, err := app.db.ListProjectsByCategory(ctx, category)
	if err != nil {
		app.serverError(w, errors.New("projects list failed"))
		return err
	}

	err = app.RenderTemplate(w, "projects/list", projects)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

func (app *Application) ListProjectsByCatAndSub(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	category := vars["category_name"]
	sub := vars["subcategory_name"]

	err := app.db.CategoryAndSubCategoryExistsPrj(ctx, category, sub)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
	}

	projects, err := app.db.ListProjectsByCategoryAndSubcategory(ctx, category, sub)
	if err != nil {
		app.serverError(w, errors.New("projects list failed"))
		return err
	}

	err = app.RenderTemplate(w, "projects/list", projects)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

func (app *Application) DeleteProject(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["project_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "invalid project id")
		return err
	}
	valid := int64(id)

	err = app.db.DeleteProject(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("project deletion failed"))
		return err
	}

	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

func (app *Application) projectUpdateForm(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["project_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "invalid project id")
		return err
	}
	valid := int32(id)

	p, err := app.db.GetProjectByID(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("project not found"))
		return err
	}

	err = app.RenderTemplate(w, "projects/update", p)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

func (app *Application) UpdateProject(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["project_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "invalid project id")
		return err
	}
	valid := int32(id)

	prj := postgres.UpdateProjectParams{
		ProjectID:   valid,
		Title:       r.FormValue("title"),
		Category:    r.FormValue("category"),
		Subcategory: r.FormValue("subcategory"),
		Description: r.FormValue("description"),
		Link:        r.FormValue("link"),
		Details:     r.FormValue("details"),
		Payments:    r.FormValue("payments"),
		FundingGoal: r.FormValue("funding_goal"),
	}
	err = app.db.UpdateProjectValidator(ctx, &prj)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.UpdateProject(ctx, &prj)
	if err != nil {
		app.serverError(w, errors.New("project update failed"))
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/categories/project/%d", valid), http.StatusFound)
	return nil
}
