package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Serve File

// Load resources
//go:embed resources
var resources embed.FS
func TestServeFile(t *testing.T)  {
	router := httprouter.New()

	// without /resources, use fs.Sub
	// ex: /files/resources/hello.txt -> /files/hello.txt
	// go in resources folder
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))
 
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	// can be print
	fmt.Println("bytes res: ", string(bytes))
	assert.Equal(t, "Hello .txt HttpRouter", string(bytes))
}

func TestServeFileGoodBye(t *testing.T)  {
	router := httprouter.New()

	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))
 
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	// can be print
	fmt.Println("bytes res: ", string(bytes))
	assert.Equal(t, "Goodbye .txt HttpRouter", string(bytes))
}
