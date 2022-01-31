package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/posts", NewPost).Methods("POST")
	myRouter.HandleFunc("/posts", GetPost).Methods("GET")
	myRouter.HandleFunc("/editposts/{id}", EditPost).Methods("PUT")
	myRouter.HandleFunc("/deletepost/{id}", DeletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	fmt.Println("Go Task")
	handlerequest()

}
