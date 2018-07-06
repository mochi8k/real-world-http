package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, _ := writer.CreateFormFile("thumbnail", "photo.jpg")
	readFile, _ := os.Open("photo.jpg")
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, _ := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	log.Println("Status: ", resp.Status)
}
