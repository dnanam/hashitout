package main

import (
	"fmt"
	"net/http"
)

// ResponseWriter is an interface and Request is a pointer
// this is the default way to handle web request in go lang
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	// and Fprint also takes in an interface for w (writer)
	fmt.Fprintf(w, "<h1>Hello Worlddss</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:ma@io.com\">ma@io.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>FAQ Page</h1><p>Do we offer support? </p>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		// 404
		//http.NotFound(w, r)
		http.Error(w, "Page not found", http.StatusNotFound)
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, "<h1>Not found</h1>")
	}
}

// Router type implements http.Handler interface from the http package
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// 404
		//http.NotFound(w, r)
		http.Error(w, "Page not found", http.StatusNotFound)
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, "<h1>Not found</h1>")

	}
}

// the http package provider http.HandlerFunc which implements http.Handler interface type
// the http package also has ServeMux type which implements http.Handler interface type

func main() {
	// setup the route
	// we know that http.HandleFunc takes in a handler which is a function and its signature matches ServeHttp from http.Handler interface
	// HandleFunc registers the handler function for the given pattern in [DefaultServeMux].
	//http.HandleFunc("/", pathHandler)
	// register the handler with a Page
	//http.HandleFunc("/contact", contactHandler)

	//var router http.HandlerFunc
	//router = pathHandler

	fmt.Println("Starting web server on port 8000")
	// ListenAndServe takes in a handler of type Handler, where Handler is an interface
	// with ServeHTTP(ResponseWriter, *Request) function. So anyone can implement the Handler interface
	// meaning I can create my own type and implement the interface
	http.ListenAndServe(":8000", http.HandlerFunc(pathHandler))
	//http.ListenAndServe(":8000", router)
}
