package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var tpl *template.Template

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.method:", r.Method)
	// if method is GET then load form, if not then upload successfull message
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "fileUpload.html", nil)
		return
	}
	// func (r *Request) ParseMultipartForm(maxMemory int64) error
	r.ParseMultipartForm(10)
	// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("fileHeader.Filename: %v\n", fileHeader.Filename)
	fmt.Printf("fileHeader.Size: %v\n", fileHeader.Size)
	fmt.Printf("fileHeader.Header: %v\n", fileHeader.Header)

	// tempFile, err := ioutil.TempFile("images", "upload-*.png")
	contentType := fileHeader.Header["Content-Type"][0]
	fmt.Println("Content Type:", contentType)
	var osFile *os.File
	// func TempFile(dir, pattern string) (f *os.File, err error)
	if contentType == "image/jpeg" {
		osFile, err = ioutil.TempFile("images", "*.jpg")
	} else if contentType == "application/pdf" {
		osFile, err = ioutil.TempFile("PDFs", "*.pdf")
	} else if contentType == "text/javascript" {
		osFile, err = ioutil.TempFile("js", "*.js")
	}
	fmt.Println("error:", err)
	defer osFile.Close()

	// func ReadAll(r io.Reader) ([]byte, error)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// func (f *File) Write(b []byte) (n int, err error)

	osFile.Write(fileBytes)

	fmt.Fprintf(w, "Your File was Successfully Uploaded!\n")
}

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":8080", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page")
}
