package main

import (
	"api/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvironment()
	paths := router.GetRouter()
	port := config.Port
	log.Printf("API LISTEN ON %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), paths))
}
