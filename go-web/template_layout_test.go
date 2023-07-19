package go_web

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Template Layout
func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	// render
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Action Layout",
		"Name": "Rocket",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println("string(body)): ", string(body))
}

// Template Layout
// func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
// 	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

// 	// render
// 	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
// 		"Title": "Template Action Range",
// 		"Hobbies": []string{
// 			"Game", "Read", "Code",
// 		},
// 	})
// }

// func TestTemplateActionRange(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
// 	recorder := httptest.NewRecorder()

// 	TemplateActionRange(recorder, request)

// 	response := recorder.Result()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	println("string(body)): ", string(body))
// }

// // Template Action With
// func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
// 	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

// 	// render
// 	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
// 		"Title": "Template Action with",
// 		"Name": "Rocket",
// 		// Case alamat present
// 		"Address" : map[string]interface{}{
// 			"Street": "Wakanda Street",
// 			"City": "Wakanda City",
// 		},

// 		// Case alamat not present
// 		// "Address" : map[string]interface{}{},
// 		//
// 	})
// }

// func TestTemplateActionWith(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
// 	recorder := httptest.NewRecorder()

// 	TemplateActionWith(recorder, request)

// 	response := recorder.Result()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	println("string(body)): ", string(body))
// }
