package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// # Download File
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")

	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	// Set Header so file will be download
	writer.Header().Add("Content-Disposition", "attachment; filename=\"" + fileName + "\"")
	http.ServeFile(writer, request, "./resources/" + fileName)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	// Browser Input
	// http://localhost:8080/?file=PZN-ICON.png

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
