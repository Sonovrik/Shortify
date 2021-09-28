package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./html/"
	}

	//http.ServeFile(w, r, http.FileServer(http.Dir(p)))
}

func main() {
	port, exists := os.LookupEnv("SHORTIFY_CONTAINER_PORT")
	if !exists {
		panic("No SHORTIFY_CONTAINER_PORT")
	}

	//fileServer := http.FileServer(http.Dir("./html"))
	//http.Handle("/html/", http.StripPrefix("/html/", fileServer))
	http.Handle("/", http.FileServer(http.Dir("./html/")))

	//http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
