package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// # Middleware

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Before handler exec
	fmt.Println("Before Execute Handler")

	// Handler exec
	middleware.Handler.ServeHTTP(writer, request)

	// After handler exec
	fmt.Println("After Execute Handler")
}

// Error Handler Middleware
type ErrorHandler struct{
	Handler http.Handler
}

// Handler
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Before handler exec
	defer func()  {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()

	// Handler exec
	errorHandler.Handler.ServeHTTP(writer, request)

	// After handler exec
}

func TestMiddleware(t *testing.T) {
	/* # Flow
		request -> server -> errorHandler -> logMiddleware -> mux
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	// Test Error
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")
		panic("Ups panic")
	})

	// When error happen in mux or logMiddleware, it can be handle by errorHandler
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
