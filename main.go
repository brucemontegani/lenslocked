package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/brucemontegani/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

// var r *chi.Mux

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func executeTemplate(w http.ResponseWriter, filePath string) {
	tpl, err := views.Parse(filePath)
	if err != nil {
		log.Print(err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h1>You have requested id: %s", id)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
