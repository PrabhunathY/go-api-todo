package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":9000"
	router := mux.NewRouter()
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Up and running")
	})
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items", addItem).Methods("POST")
	router.HandleFunc("/items", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", removeItem).Methods("DELETE")
	log.Println("Server listening on port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalln(err)
	}
}
