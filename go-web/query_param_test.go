package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Query Parameter
// handler func
func SayHello(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T)  {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=Rocket", nil)
	recorder := httptest.NewRecorder()

	// handler call
	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("string(body): ", string(body))
}

// Multiple Query Parameter
func MultipleQueryParam(writer http.ResponseWriter, request *http.Request)  {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParam(t *testing.T)  {
	request := httptest.NewRequest("GET", "http://localhost:8080?first_name=Rocket&last_name=Racoon", nil)
	recorder := httptest.NewRecorder()

	// handler call
	MultipleQueryParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("string(body): ", string(body))
}

// Multiple Value Query Parameter
func MultipleParamValue(writer http.ResponseWriter, request *http.Request)  {
	queryParam := request.URL.Query()
	names := queryParam["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParamValues(t *testing.T)  {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=Rocket&name=Racoon&name=tasty", nil)
	recorder := httptest.NewRecorder()

	// handler call
	MultipleParamValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("string(body): ", string(body))
}