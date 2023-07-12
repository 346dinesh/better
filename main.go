package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("running")
	http.HandleFunc("/myhandler", myHandler)
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// Execute your API calls here
	fmt.Println("Button clicked!")
	// You can make API calls or perform any other operations
}
