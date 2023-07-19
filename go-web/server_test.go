package go_web

import (
	"net/http"
	"testing"
)

// Server
func TestServer(t *testing.T)  {
	server := &http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}