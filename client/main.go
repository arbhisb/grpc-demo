package main

import (
	"clientgrpc/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test/{message}", handler.SayHello).Methods("GET")
	r.HandleFunc("/testpost", handler.SayHelloPost).Methods("POST")
	http.ListenAndServe(":3000", r)
}
