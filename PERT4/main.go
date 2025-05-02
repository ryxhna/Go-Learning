package main

import (
	"PERT4/handler"
	"log"
	"net/http"
)

// TODO: ganti ke dua digit terakhir NPM
func main() {
	http.HandleFunc("/api/", handler.API)
	log.Println("localhost : 8081")
	http.ListenAndServe(":8081", nil)
}
