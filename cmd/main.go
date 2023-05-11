package main

import (
	"Tugas3EAI/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/graphql", handler.GraphqlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
