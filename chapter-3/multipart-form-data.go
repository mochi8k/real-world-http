package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"net/textproto"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, _ := writer.CreatePart(part)

	// fileWriter, _ := writer.CreateFormFile("thumbnail", "photo.jpg")
	readFile, _ := os.Open("photo.jpg")
	defer readFile.Close()
	io.Copy(fileWriter, readFile)

	writer.Close()

	resp, _ := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	log.Println("Status: ", resp.Status)
}
