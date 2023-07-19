package go_web

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Template Data Map
func TemplateDataMap(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name": "Rocket",
		"Address": map[string]interface{}{
			"Street": "Halfworld Street",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}

// Template with Struct
type Address struct {
	Street	string
}

type Page struct {
	Title	string
	Name 	string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Rocket",
		Address: Address{
			Street: "Halfworld Street",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}
