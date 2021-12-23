package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp, "Up and running.....")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening the port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
