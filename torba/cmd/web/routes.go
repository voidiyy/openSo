package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type funky func(w http.ResponseWriter, r *http.Request) error

func (app *application) handleGET(f funky) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("error handle %v GET request, %v", err, r.RequestURI)
		}
	}
}

func (app *application) handlePOST(f funky) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("error handle %v POST request, %v", err, r.RequestURI)
		}
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, templateName string, data map[string]interface{}) {
	t, err := template.ParseFiles(templateName)
	if err != nil {
		app.errorL.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		app.errorL.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	//index pages (GET)
	router.HandleFunc("/", app.handleGET(app.home)).Methods("GET")
	router.HandleFunc("/about", app.handleGET(app.about)).Methods("GET")
	router.HandleFunc("/contact", app.handleGET(app.contact)).Methods("GET")

	//[[][][][][][][][][][][][][][][[][][][]][][][]][][][][][][][][][][][][][][][][][]

	//user routes
	router.HandleFunc("/user/profile/{user_id}", nil).Methods("GET")
	router.HandleFunc("/user/profile/{user_id}/update", nil).Methods("PUT")
	router.HandleFunc("/user/profile/{user_id}/delete", nil).Methods("DELETE")
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
	router.HandleFunc("/author/{author_id}/profile", nil).Methods("GET")
	router.HandleFunc("/author/{author_id}/profile/update", nil).Methods("PUT")
	router.HandleFunc("/author/{author_id}/profile/delete", nil).Methods("DELETE")
	router.HandleFunc("/author/{author_id}/profile/logout", nil).Methods("POST")
	router.HandleFunc("/author/{author_id}/profile/stats", nil).Methods("POST")

	//author projects
	router.HandleFunc("/author/{author_id}/profile/projects", nil).Methods("GET")
	router.HandleFunc("/author/{author_id}/profile/projects/create", nil).Methods("POST")
	router.HandleFunc("/author/{author_id}/profile/projects/{project_id/update}", nil).Methods("PUT")
	router.HandleFunc("/author/{author_id}/profile/projects/{project_id/delete}", nil).Methods("DELETE")

	//api and form
	router.HandleFunc("/author/create", nil).Methods("GET")
	router.HandleFunc("/author/create", nil).Methods("POST")

	//api and form
	router.HandleFunc("/author/login", nil).Methods("GET")
	router.HandleFunc("/author/login", nil).Methods("POST")

	//[][][][][][][][][][][][][][][][]][][][][]][][][][][][][][][][][][]][][][][][][][][][]

	//categories project routes
	router.HandleFunc("/categories", nil).Methods("GET")

	router.HandleFunc("/categories/project", nil).Methods("GET")
	router.HandleFunc("/categories/project/{category_name}", nil).Methods("GET")
	router.HandleFunc("/categories/project/{category_name}/{subcategory_name}", nil).Methods("GET")
	router.HandleFunc("/categories/project/{category_name}/{subcategory_name}/{project_id}", nil).Methods("GET")
	router.HandleFunc("/categories/project/{category_name}/{subcategory_name}/{project_id}/donate", nil).Methods("GET")
	router.HandleFunc("/categories/project/{category_name}/{subcategory_name}/{project_id}/comments", nil).Methods("GET")

	//categories org routes
	router.HandleFunc("/categories/organization", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{category_name}", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{category_name}/{subcategory_name}", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{category_name}/{subcategory_name}/{organization_id}", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{category_name}/{subcategory_name}/{organization_id}/donate", nil).Methods("GET")
	router.HandleFunc("/categories/organization/{category_name}/{subcategory_name}/{organization_id}/comments", nil).Methods("GET")

	//organizations routes

	return router
}
