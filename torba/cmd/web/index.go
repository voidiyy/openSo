package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *Application) about(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	if r.URL.Path != "/about" {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("page not found")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("../../ui/index/about.gohtml")
	if err != nil {
		http.NotFound(w, r)
		app.errorL.Printf("template parsing error: %v", err)
		return fmt.Errorf("template parsing error")
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.NotFound(w, r)
		app.errorL.Printf("template executing error: %v", err)
		return fmt.Errorf("template executing error")
	}

	return nil
}

func (app *Application) contact(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	if r.URL.Path != "/contact" {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("page not found")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	templ, err := template.ParseFiles("../../ui/index/contact.gohtml")
	if err != nil {
		http.NotFound(w, r)
		app.errorL.Printf("template parsing error: %v", err)
		return fmt.Errorf("template parsing error")
	}

	err = templ.Execute(w, nil)
	if err != nil {
		http.NotFound(w, r)
		app.errorL.Printf("template executing error: %v", err)
		return fmt.Errorf("template executing error")
	}

	return nil
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("page not found")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	templ, err := template.ParseFiles("../../ui/index/home.gohtml")
	if err != nil {
		http.NotFound(w, r)
		app.errorL.Printf("template parsing error: %v", err)
		return fmt.Errorf("template parsing error")
	}
	err = templ.Execute(w, nil)
	if err != nil {
		http.NotFound(w, r)
		app.errorL.Printf("template executing error: %v", err)
		return fmt.Errorf("template executing error")
	}

	return nil
}
