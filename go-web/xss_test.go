package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// XSS Cross Site Scripting

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	// render
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		// Scenario 1: use html/template
		// not render as js
		// "Body": "<p>Ini Adalah Body</p>",
		
		// Scenario 2: use html/template
		// not render as js
		"Body": "<p>Ini Adalah Body<script>alert('Anda di Hack')</script></p>",
	})

	// Escape Response:
	// &lt;p&gt;Ini Adalah Body&lt;/p&gt;
}

func TestTemplateAutoEscape_(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Disable Auto Escape
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	// render
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		// Scenario 1: use html/template
		// Disable auto escape, render as HTML
		"Body": template.HTML("<h1>Disabled Auto Escape HTML!</h1>"),

		// Disable Escape Response:
		// Disabled Auto Escape HTML!
	})
}

func TestTemplateAutoEscapeDisabled_(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Contoh CSS
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	// render
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template XSS",
		"Body": template.HTML(request.URL.Query().Get("body")),

		// Disable Escape Response:
		// data from body
	})
}

func TestTemplateXSS_(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	/*
	// browser input scenario:
		- http://localhost:8080/?body=<script>alert("Anda di Hack")</script>
	*/

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
