package main

import (
	"context"
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9808", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	res, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from client!"})
	if err != nil {
		log.Fatalf("Error while calling SayHello(): %s", err)
	}

	log.Printf("Received message from server :", res.Body)
}
