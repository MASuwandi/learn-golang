package go_web

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Template with Text / string
func SimpleHTML(writer http.ResponseWriter, request *http.Request)  {
	// create template
	templateText := `<html><body>{{.}}</body></html>`

	// parsing template
	t, err := template.New("SIMPLE").Parse(templateText)

	// alternative parsing, no need to check err because must already handle it
	// t := template.Must(template.New("SIMPLE").Parse(templateText))

	if err != nil {
		panic(err)
	}
	
	// write template to response
	// dynamic data will replace .
	// t.ExecuteTemplate(writer, name, data)
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template!")
}

func TestHTMLTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}

// Template with File
func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request)  {
	// parsing template from file location
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template File!")
}

func TestHTMLTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}

// Template with Directory
func SimpleDir(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Directory!")
}

func TestTemplateDir(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleDir(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}

// Template with Golang Embed (Recommended)

// embed from template caching


func TemplateEmbed(writer http.ResponseWriter, request *http.Request)  {
	// without dot because from embed
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Embed!")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}