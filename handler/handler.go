package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/graphql-go/graphql"
	"Tugas3EAI/graph"
)

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Handle GET request
		GetBooks(w, r)
	case http.MethodPost:
		// Handle POST request
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Execute the GraphQL query
		params := graphql.Params{
			Schema:        graph.Schema,
			RequestString: string(body),
		}
		result := graphql.Do(params)

		// Write the response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(result)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Get list of books
	books, err := graph.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(books)
}