package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Handle defalt route
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Construct address
	addr := ":" + port

	// Start http server
	log.Println(fmt.Sprintf(`Server is listening on address "%s"`, addr))
	http.ListenAndServe(addr, nil)
}
