package main

import (
	"fmt"
	"net/http"
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

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		http.Error(w, "Siden blev ikke fundet", http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "Page not found")
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Siden blev ikke fundet", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)

	/* http.HandleFunc("/", homeHandler)
	// http.Handle("/contact", http.HandlerFunc(contactHandler))


	// http.Handle() - Interface with the ServeHTTP method
	// http.Handle("/", http.Handler)
	// http.handlerFunc() - A function type that accepts same args as ServeHTTP method. also implements http.Handler
	 http.HandleFunc("/", pathHandler) */
}
