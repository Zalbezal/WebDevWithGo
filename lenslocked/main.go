package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/zalbezal/lenslocked/controllers"
	"github.com/zalbezal/lenslocked/templates"
	"github.com/zalbezal/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.Get("/static", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "static.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)

}
