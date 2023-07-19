package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// Handler
func TestHandler(t *testing.T) {
	// Handler
	var handler http.HandlerFunc = func(writer http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello Web")
	}

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Serve Mux
func TestServeMux(t *testing.T) {
	// router
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello Web")
	})

	mux.HandleFunc("/user", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello /user")
	})

	mux.HandleFunc("/user/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello /user/")
	})

	mux.HandleFunc("/user/image", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello /user/image")
	})

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Request
func TestRequest(t *testing.T) {
	// router
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello Web")
		fmt.Fprintln(writer, r.Method)
		fmt.Fprintln(writer, r.RequestURI)
	})

	mux.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello /user")
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	})

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
