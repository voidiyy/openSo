package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"openSo/internal/postgres"
	"strconv"
	"strings"
)

//categories/project/{project_id}

func (app *Application) organizationByID(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid organization ID")
		return err
	}
	valid := int64(id)

	org, err := app.db.GetOrganizationByID(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	err = app.RenderTemplate(w, "", org)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

//categories/project/create GET

func (app *Application) organizationCreateForm(w http.ResponseWriter, r *http.Request) error {

	err := app.RenderTemplate(w, "", nil)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

//categories/project/create POST

func (app *Application) organizationCreate(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	org := postgres.CreateOrgParams{
		AuthorID:       1,
		Category:       r.FormValue("category"),
		Subcategory:    r.FormValue("subcategory"),
		Name:           r.FormValue("name"),
		Description:    r.FormValue("description"),
		Website:        r.FormValue("website"),
		ContactEmail:   r.FormValue("email"),
		AdditionalInfo: r.FormValue("additional_info"),
	}

	err := app.db.CreateOrgValidator(ctx, &org)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.CreateOrganization(ctx, &org)
	if err != nil {
		app.serverError(w, errors.New("error creating organization"))
	}

	red, err := app.db.GetOrganizationByName(ctx, org.Name)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/categories/organization/%v", red.OrgID), http.StatusSeeOther)
	return nil
}

//categories/organizations/{org_id}/delete

func (app *Application) DeleteOrganization(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["org_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid organization ID")
		return err
	}
	valid := int64(id)

	err = app.db.DeleteOrganization(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

// categories/organizations/{org_id}/update GET
func (app *Application) organizationUpdateForm(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["org_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid organization ID")
		return err
	}
	valid := int64(id)

	o, err := app.db.GetOrganizationByID(ctx, valid)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	err = app.RenderTemplate(w, "", o)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}
	return nil
}

// categories/organizations/{org_id}/update GET

func (app *Application) organizationUpdate(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["org_id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid organization ID")
		return err
	}
	valid := int64(id)

	org := postgres.UpdateOrgParams{
		OrgID:          valid,
		Category:       r.FormValue("category"),
		Subcategory:    r.FormValue("subcategory"),
		Name:           r.FormValue("name"),
		Description:    r.FormValue("description"),
		Website:        r.FormValue("website"),
		ContactEmail:   r.FormValue("email"),
		AdditionalInfo: r.FormValue("additional_info"),
	}

	err = app.db.UpdateOrgValidator(ctx, &org)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return err
	}

	err = app.db.UpdateOrg(ctx, &org)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/categories/organization/%v", valid), http.StatusSeeOther)
	return nil
}

//categories/organization/{category_name}

func (app *Application) ListOrgByCategories(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)

	category := strings.TrimSpace(vars["category_name"])

	err := app.db.CategoryExistsOrg(ctx, category)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	organizations, err := app.db.ListOrganizationsByCategory(ctx, category)
	if err != nil {
		app.serverError(w, errors.New("organizations does not exist"))
		return err
	}

	err = app.RenderTemplate(w, "", organizations)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

//categories/organization/{category_name}/{subcategory}

func (app *Application) ListOrgByCatSub(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	category := strings.TrimSpace(vars["category_name"])
	subcategory := strings.TrimSpace(vars["subcategory_name"])

	err := app.db.CategoryAndSubCategoryExistsOrg(ctx, category, subcategory)
	if err != nil {
		app.serverError(w, errors.New("organization does not exist"))
		return err
	}

	organizations, err := app.db.ListOrganizationsByCatSub(ctx, category, subcategory)
	if err != nil {
		app.serverError(w, errors.New("organizations does not exist"))
		return err
	}

	err = app.RenderTemplate(w, "", organizations)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}

func (app *Application) ListOrgByAuthor(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	vars := mux.Vars(r)
	authorID, err := strconv.ParseInt(vars["author_id"], 10, 64)
	if err != nil {
	}

	_, err = app.db.GetAuthorByID(ctx, authorID)
	if err != nil {
		app.serverError(w, errors.New("author does not exist"))
		return err
	}

	organizations, err := app.db.ListOrganizationsByAuthor(ctx, authorID)
	if err != nil {
		app.serverError(w, errors.New("organizations does not exist"))
		return err
	}

	err = app.RenderTemplate(w, "", organizations)
	if err != nil {
		app.serverError(w, errors.New("error rendering template"))
		return err
	}

	return nil
}
