package main

import (
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	mux := http.NewServeMux()

	http.HandleFunc("/hello", yourFunction)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":4010", nil))
}

func yourFunction(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, Aleksei")

}
