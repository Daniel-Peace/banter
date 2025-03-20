package main

import (
	"fmt"
	"log"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutEndPoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"text\":\"Home Page\"")
}

func aboutEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About me")
}

func main() {
	fmt.Println("Hello world")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
