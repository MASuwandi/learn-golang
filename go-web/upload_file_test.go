package go_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// # Upload
func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload_form.gohtml", nil)
}

// Handler
func Upload(writer http.ResponseWriter, request *http.Request) {
	// Parsing
	// request.ParseMultipartForm(100 << 20) // max 100 mb
	// Behind it will parse first, def max 32mb
	// then Take the file
	// param base on name from:
	// 		<input type="file" name="file">
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	// Save file and create Dir
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// Move uploaded file to file destination
	// return writed size and err
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	// Take non file
	name := request.PostFormValue("name")

	// Render into template
	myTemplates.ExecuteTemplate(writer, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm_(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Unit Test simulate upload file

//go:embed resources/PZN-ICON.png
var uploadFileTest []byte

func TestUploadFile_(t *testing.T) {

	// Tempat menyimpan binary body
	body := new(bytes.Buffer)

	// Set format as multipart
	writer := multipart.NewWriter(body)

	// # Input data
	// Input Text
	writer.WriteField("name", "Rocket")

	// Create File
	// param base on name from:
	// 		<input type="file" name="file">
	// (name, writed file name)
	file, _ := writer.CreateFormFile("file", "contoh_upload.png")

	// Fill in the Creted File
	file.Write(uploadFileTest)

	// Close writer to. handle use memory
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	
	// Set Content Type
	// writer.FormDataContentType() == "multipart/form-data"
	request.Header.Set("Content-Type", writer.FormDataContentType())
	
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyRes, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(bodyRes))
}
