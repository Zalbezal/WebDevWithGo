package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jeppe-christiansen@hotmail.com\">jeppe-christiansen@hotmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
	fmt.Fprint(w, `<ul>
		<li><b>Q: is this a course</b> A: No this is a test site</li>
		
		<li><b>Q: Is it free</b> A: well yeah if you can find it</li>
		<li><b>How do I contact support?</b> Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a></li>
		</ul>`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

	/* http.HandleFunc("/", homeHandler)
	// http.Handle("/contact", http.HandlerFunc(contactHandler))


	// http.Handle() - Interface with the ServeHTTP method
	// http.Handle("/", http.Handler)
	// http.handlerFunc() - A function type that accepts same args as ServeHTTP method. also implements http.Handler
	 http.HandleFunc("/", pathHandler) */
}
