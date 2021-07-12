package handler

import (
	"clientgrpc/api"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type sayHelloRequest struct {
	Message string
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["message"]
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: message})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}

func SayHelloPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t sayHelloRequest
	errdecode := decoder.Decode(&t)
	if errdecode != nil {
		panic(errdecode)
	}
	log.Println(t.Message)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: t.Message})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
