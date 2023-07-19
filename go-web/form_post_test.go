package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Form Post
func FormPost(writer http.ResponseWriter, request *http.Request)  {
	// manual
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")

	// Better alternative
	// request.PostFormValue()

	fmt.Fprintf(writer, "Hi Body, %s %s", first_name, last_name)

}

func TestFormPost(t *testing.T)  {
	requestBody := strings.NewReader(("first_name=Rocket&last_name=Racoon"))
	request := httptest.NewRequest("POST", "http://localhost:8080?name=Rocket", requestBody)
	// standar header for post body
	// wrong type can lead to unreadable body data
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	// handler call
	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("string(body): ", string(body))
}
