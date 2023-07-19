package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Cookie
// handler
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// define value for name / key, value, path / url
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	// set cookie via query param value
	cookie.Value = request.URL.Query().Get("name")
	// which url cookie active
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	// define value for name / key, value, path / url
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// Cookie with Server
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("Test Cookie Done")
}

// Test Set Cookie
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
	

	fmt.Println("Test Set Cookie Done")
}

// Test Client Request with Cookie / Mengambil Cookie
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Rocket"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	
	fmt.Println("string(body): ", string(body))
}
