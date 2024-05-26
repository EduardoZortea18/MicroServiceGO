package main

import (
	"log"
	"net/http"

	common "github.com/EduardoZortea18/common"
	pb "github.com/EduardoZortea18/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(
		insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	log.Println("Dialing order service ", orderServiceAddr)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting http server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
