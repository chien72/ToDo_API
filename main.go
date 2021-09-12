package main

import (
	"ToDo/server/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")
	fmt.Println("hello")
	log.Fatal(http.ListenAndServe(":8080", r))
}
