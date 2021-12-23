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
		fmt.Println(resp, "Up and running.....")
	})
	log.Println("Server listening the port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
