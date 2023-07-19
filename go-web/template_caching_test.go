package go_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"html/template"
	// "text/template"
)

// Template Caching

//go:embed templates/*.gohtml
var templates embed.FS

// Set as global so only parse once
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	// render
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching!")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}
