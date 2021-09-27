package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!!")
}

func main() {
	port, exists := os.LookupEnv("SHORTIFY_CONTAINER_PORT")
	if !exists {
		panic("No SHORTIFY_CONTAINER_PORT")
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
