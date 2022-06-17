package main

import (
	"fmt"
	"gotest/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Running server on port 1111...")
	log.Fatal(http.ListenAndServe(":1111", r))
}