package main

import (
	"fmt"
	"net/http"
)

func (app *application) userForm(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/user/create" {
		http.NotFound(w, r)
		app.errorL.Printf("invalid path: %s", r.URL.Path)
		return fmt.Errorf("invalid path")
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		app.errorL.Printf("Method not allowed: %s", r.Method)
		return fmt.Errorf("method not allowed")
	}

	app.render(w, r, "../../ui/user.html", nil)
	return nil
}

func (app *application) about(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	if r.URL.Path != "/about" {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("page not found")
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "AboutPage!")
	return nil
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	if r.URL.Path != "/contact" {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("page not found")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ContactPage!")

	return nil
}

func (app *application) home(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("page not found")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Home Page!")

	return nil
}
