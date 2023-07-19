package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Header Request / Client
func RequstHeader(writer http.ResponseWriter, request *http.Request)  {
	contentType := request.Header.Get("content-type")

	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T)  {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequstHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("string(body): ", string(body))
}

// Header Response
func ResponseHeader(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Add("X-Powered-By", "Programmer Zaman now")

	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T)  {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("string(body): ", string(body))

	fmt.Println("response header: ", response.Header.Get("x-powered-by"))

}
