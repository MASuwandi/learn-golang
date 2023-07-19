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

// Test HTTP Router
func TestRouter(t *testing.T)  {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)  {
		fmt.Fprint(writer, "Hello Get Rocket!")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	// can be print
	fmt.Println("bytes res: ", string(bytes))
	assert.Equal(t, "Hello Get Rocket!", string(bytes))
}