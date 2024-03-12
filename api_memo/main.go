package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func messageHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello, this is the message from the API!")
	case "POST":
		var newPost Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Printf("%v\r", newPost)
		}
	}

}

func main() {
	http.HandleFunc("/", messageHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Printf("Server Started")
}
