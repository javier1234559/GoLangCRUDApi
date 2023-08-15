package main

import (
	"fmt"
	"log"
	"net/http"
)

type Awesomizer interface {
	Awesomize() string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	fmt.Println("Server is running !!")
	handleRequests()
}
