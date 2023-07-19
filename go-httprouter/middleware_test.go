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

// Middleware

// Struct Middleware
type LogMiddleware struct{
	// Handler http.Handler
	http.Handler
}

// Middleware for log
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	fmt.Println("Receive request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T)  {
	// request -> server -> middleware -> router
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)  {
		fmt.Fprint(writer, "Middleware")
	})

	middleware := LogMiddleware{Handler: router}
	// Shortcut
	// middleware := LogMiddleware{router}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println("body res: ", string(body))
	assert.Equal(t, "Middleware", string(body))
}
