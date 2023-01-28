package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadConfiguration()

	fmt.Println("Server is running on port", config.Port)

	fmt.Println("rodando api")

	r := router.GenerateRoutes()

	log.Fatal(http.ListenAndServe(":5000", r))
}
