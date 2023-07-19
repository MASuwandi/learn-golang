package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Template Function

type MyPage struct{
	Name	string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// Invoke func and send "Rocket" as param
	// call to data struct from render
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Rocket"}}`))

	// render
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Thor",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}

// Template Function Global
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	// render
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Thor",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}

// Template Menambah Function Global
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	// Create Template
	t := template.New("FUNCTION")

	// Create Function name "upper"
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// Parsing
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	// Render
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Thor",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}

// Template Function Pipelines
func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	// Create Template
	t := template.New("FUNCTION")

	// Create Function name "upper"
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"sayHello": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// Parsing
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	// Render
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Thor",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("string(body)): ", string(body))
}