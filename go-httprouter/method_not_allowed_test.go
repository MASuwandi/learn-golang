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

// Method Not Allowed Handler

// Return "Method Not Allowed" in beowser
func TestMethodNotAllowedHandler(t *testing.T)  {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Tak Boleh")
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		fmt.Fprint(writer, "POST")
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println("body res: ", string(body))
	assert.Equal(t, "Tak Boleh", string(body))
}
