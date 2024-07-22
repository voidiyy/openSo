package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
)

type funky func(w http.ResponseWriter, r *http.Request) error

func (app *Application) handleGET(f funky) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		if err := f(w, r); err != nil {
			app.errorL.Printf("error handling this GET path %v :: %v", r.RequestURI, err)
		} else {
			app.infoL.Printf("succesfuly handled GET path %v", r.RequestURI)
		}
	}
}

func (app *Application) handlePOST(f funky) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		if err := f(w, r); err != nil {
			app.errorL.Printf("error handling this POST path %v :: %v", r.RequestURI, err)
		} else {
			app.infoL.Printf("succesfuly handled POST path %v", r.RequestURI)
		}
	}
}

func (app *Application) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		app.errorL.Printf("error parsing template: %v", err)
		return err
	}
	err = t.Execute(w, data)
	if err != nil {
		app.errorL.Printf("error executing template: %v", err)
		return err
	}
	return nil
}

func (app *Application) Templates(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) error {
	t, err := template.ParseGlob(filepath.Join("../../ui/", "*", "*.gohtml"))
	if err != nil {
		app.errorL.Printf("error parsing template patern: %v", err)
	}
	err = t.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		app.errorL.Printf("error executing template: %v", err)
	}
	return nil
}

func (app *Application) clientError(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)

	app.errorL.Printf("Client error - Status: %d, Message: %s", status, message)
}

func (app *Application) serverError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)

	app.errorL.Printf("Server error - Status: %d, Error: %v ", http.StatusInternalServerError, err)
}

func (app *Application) routes() *mux.Router {
	router := mux.NewRouter()

	//index pages (GET)
	router.HandleFunc("/", app.handleGET(app.home)).Methods("GET")
	router.HandleFunc("/about", app.handleGET(app.about)).Methods("GET")
	router.HandleFunc("/contact", app.handleGET(app.contact)).Methods("GET")

	//[[][][][][][][][][][][][][][][[][][][]][][][]][][][][][][][][][][][][][][][][][]

	router.HandleFunc("/nimda", app.handleGET(app.panel)).Methods("GET")
	router.HandleFunc("/nimda/users", app.handleGET(app.listUsers)).Methods("GET")
	router.HandleFunc("/nimda/authors", app.handleGET(app.listAuthors)).Methods("GET")
	router.HandleFunc("/nimda/projects", app.handleGET(app.listUsers)).Methods("GET")
	router.HandleFunc("/nimda/organizations", app.handleGET(app.listUsers)).Methods("GET")

	//user routes
	router.HandleFunc("/user/profile/{user_id}", app.handleGET(app.UserProfile)).Methods("GET")

	router.HandleFunc("/user/profile/{user_id}/update", app.handlePOST(app.UserUpdateProfile)).Methods("POST")
	router.HandleFunc("/user/profile/{user_id}/update", app.handleGET(app.userUpdatePage)).Methods("GET")

	router.HandleFunc("/user/profile/{user_id}/delete", app.handlePOST(app.UserDeleteProfile)).Methods("POST")
	router.HandleFunc("/user/profile/{user_id}/logout", nil).Methods("POST")

	//this to may be implemented into profile like sub body page
	router.HandleFunc("/user/profile/{user_id}/stats", nil).Methods("GET")
	router.HandleFunc("/user/profile/{user_id}/supported", nil).Methods("GET")

	//api and form
	router.HandleFunc("/user/create", app.handleGET(app.userForm)).Methods("GET")
	router.HandleFunc("/user/create", app.handlePOST(app.CreateUser)).Methods("POST")

	//api and form
	router.HandleFunc("/user/login", nil).Methods("GET")
	router.HandleFunc("/user/login", nil).Methods("POST")

	//[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]

	//author routes
	router.HandleFunc("/author/{author_id}/profile", app.handleGET(app.ProfileAuthor)).Methods("GET")

	router.HandleFunc("/author/{author_id}/profile/update", app.handleGET(app.updateProfileAuthor)).Methods("GET")
	router.HandleFunc("/author/{author_id}/profile/update", app.handlePOST(app.UpdateProfileAuthor)).Methods("POST")

	router.HandleFunc("/author/{author_id}/profile/delete", app.handlePOST(app.DeleteAuthor)).Methods("DELETE")
	router.HandleFunc("/author/{author_id}/profile/logout", nil).Methods("POST")
	router.HandleFunc("/author/{author_id}/profile/stats", nil).Methods("POST")

	//author projects
	router.HandleFunc("/author/{author_id}/profile/projects", nil).Methods("GET")
	router.HandleFunc("/author/{author_id}/profile/projects/create", nil).Methods("POST")
	router.HandleFunc("/author/{author_id}/profile/projects/{project_id/update}", nil).Methods("PUT")
	router.HandleFunc("/author/{author_id}/profile/projects/{project_id/delete}", nil).Methods("DELETE")

	//api and form
	router.HandleFunc("/author/create", app.handleGET(app.authorForm)).Methods("GET")
	router.HandleFunc("/author/create", app.handlePOST(app.CreateFullAuthor)).Methods("POST")

	//api and form
	router.HandleFunc("/author/login", nil).Methods("GET")
	router.HandleFunc("/author/login", nil).Methods("POST")

	//[][][][][][][][][][][][][][][][]][][][][]][][][][][][][][][][][][]][][][][][][][][][]

	//categories project routes
	router.HandleFunc("/categories", nil).Methods("GET")

	router.HandleFunc("/categories/project/{project_id}", app.handleGET(app.projectByID)).Methods("GET")
	router.HandleFunc("/categories/project/{project_id}/delete", nil).Methods("POST")

	router.HandleFunc("/categories/project/{project_id}/update", nil).Methods("POST")
	router.HandleFunc("/categories/project/{project_id}/update", nil).Methods("GET")

	router.HandleFunc("/categories/project/{project_id}/donate", nil).Methods("GET")
	router.HandleFunc("/categories/project/{project_id}/comments", nil).Methods("GET")

	router.HandleFunc("/categories/project/create", app.handlePOST(app.CreateProject)).Methods("POST")
	router.HandleFunc("/categories/project/create", app.handleGET(app.CreateProjectForm)).Methods("GET")

	router.HandleFunc("/categories/project/{category_name}", app.handleGET(app.ListProjectsByCategory)).Methods("GET")
	router.HandleFunc("/categories/project/{category_name}/{subcategory_name}", app.handleGET(app.ListProjectsByCatAndSub)).Methods("GET")

	//categories org routes
	router.HandleFunc("/categories/organization/author/{author_id}", nil).Methods("GET")

	router.HandleFunc("/categories/organization/{org_id}", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{org_id}/delete", nil).Methods("POST")

	router.HandleFunc("/categories/organization/{org_id}/update", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{org_id}/update", nil).Methods("POST")

	router.HandleFunc("/categories/organization/{organization_id}/donate", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{organization_id}/comments", nil).Methods("GET")

	router.HandleFunc("/categories/organization/create", nil).Methods("POST")
	router.HandleFunc("/categories/organization/create", nil).Methods("GET")

	router.HandleFunc("/categories/organization/{category_name}", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{category_name}/{subcategory_name}", nil).Methods("GET")

	//organizations.sql routes

	return router
}
