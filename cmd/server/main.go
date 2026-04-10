package main

import (
	"fmt"
	"log"
	"net/http"

	_ "go-link-storage-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title go-link-storage-service API
// @version 1.0
// @description API for the go-link-storage-service application
// @host localhost:4010
// @BasePath /
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", yourFunction)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":4010", mux))
}

// @Summary Say hello
// @Description Returns a greeting message
// @Tags example
// @Produce plain
// @Success 200 {string} string "Hello, Aleksei"
// @Router /hello [get]
func yourFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Aleksei")
}
