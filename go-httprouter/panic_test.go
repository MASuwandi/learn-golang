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

// Panic Handler

// Panic Case Without Handler
// Return "ERR_EMPTY_RESPONSE" in beowser
func TestPanic_(t *testing.T)  {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
		panic("Ups")
	})

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}

// Panic Case With Handler
func TestPanicHandler(t *testing.T)  {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{})  {
		fmt.Fprint(writer, "Panic: ", error)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)  {
		panic("Ups")
	})
 
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	// can be print
	fmt.Println("body res: ", string(body))
	assert.Equal(t, "Panic: Ups", string(body))
}
