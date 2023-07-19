package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)



func HelloHandLer(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello Web")
}

func TestHttp(t *testing.T)  {
	// http request dan recorder
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	// call handler
	HelloHandLer(recorder, request)

	// check response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("bodyString: ", bodyString)
}