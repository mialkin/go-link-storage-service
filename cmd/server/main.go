package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", yourFunction)

	log.Fatal(http.ListenAndServe(":4010", nil))
}

func yourFunction(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, Aleksei")

}
