package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, _ := os.Open("main-get.go")
	log.Println(file)
	resp, _ := http.Post("http://localhost:18888", "text/plain", file)
	log.Println("Status: ", resp.Status)
}
