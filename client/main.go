package main

import (
	pb "grpc/proto"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("SERVER_PORT")
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
