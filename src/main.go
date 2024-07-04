// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	client := &DefaultRPCClient{}
	router := mux.NewRouter()
	router.HandleFunc("/getBlockNumber", getBlockNumber(client)).Methods("POST")
	router.HandleFunc("/getBlockByNumber", getBlockByNumber(client)).Methods("POST")
	router.HandleFunc("/health_check", healthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
