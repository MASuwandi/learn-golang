package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Not Found Handler

// Not Found Case Without Handler
// Return "404 page not found" in beowser
func TestNotFound_(t *testing.T)  {
	router := httprouter.New()

	// Test:
	// Hit endpoint

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}

// Panic Case With Handler
func TestNotFoundHandler_(t *testing.T)  {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Tak Dapat")
	})

	// Test:
	// Hit endpoint

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}

func TestNotFoundHandlerPZN(t *testing.T)  {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Tak Dapat")
	})

	// Test:
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println("body res: ", string(body))
	assert.Equal(t, "Tak Dapat", string(body))
}
