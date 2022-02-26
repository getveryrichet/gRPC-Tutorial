package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/getveryrichet/gRPC-Tutorial/chat"
)

func main() {

	var conn *grpc.ClientConn
	// connection to server
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// new client
	c := chat.NewChatServiceClient(conn)

	// function call
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!", Items: []string{"item1", "item2"}})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	response, err = c.BroadcastMessage(context.Background(), &chat.Message{Body: "Message to Broadcast!"})
	if err != nil {
		log.Fatalf("Error when calling Broadcast Message: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
	log.Printf("Response from Richet's gRPC server: %s, %s", response.Body, response.Items)

}
