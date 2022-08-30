package main

import (
	"all-my-gear-backend-go/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
    r := router.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
    fmt.Println("Starting server on the port 8080...")

    log.Fatal(http.ListenAndServe(":8080", r))
}