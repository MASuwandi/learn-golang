package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// Serve File
// dynamic file selection
func ServeFile(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Serve File with Golang Embed
//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
