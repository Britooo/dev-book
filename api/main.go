package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	paths := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", paths))
}
