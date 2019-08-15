package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		log.Fatal("POD_NAME environment variable is required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "request served by pod %s\n", podName)
	})

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
