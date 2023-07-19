
package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Response Code
func ResponseCode(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Name is empty")
	} else {
		fmt.Fprintf(writer, "Hello name is %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T)  {
	// fail scenarion
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)

	// content-type
	// request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	// handler call
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Status: ", response.Status)
	fmt.Println("string(body): ", string(body))
}

func TestResponseCodeValid(t *testing.T)  {
	// success scenarion
	request := httptest.NewRequest("GET", "http://localhost:8080?name=Rocket", nil)

	recorder := httptest.NewRecorder()

	// handler call
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Status: ", response.Status)
	fmt.Println("string(body): ", string(body))
}
