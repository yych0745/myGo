package main

import (
	lissa "awesomeProject/slice"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissa.Lissajous(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}
