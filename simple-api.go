package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Simple GET handler
func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is GET request")
}

// Struct for POST JSON
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// POST handler
func postHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Received user: %s, Age: %d", user.Name, user.Age)
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
