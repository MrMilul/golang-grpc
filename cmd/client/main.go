package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "example.com/laptop_store/proto"
	"example.com/laptop_store/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	// Get address
	serverAddress := flag.String("address", "", "The server address")
	flag.Parse()
	log.Printf("Dial Server %s", *serverAddress)

	// laptop sample
	laptop := sample.NewLaptop()

	// Dial to the server
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial server", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	// call the CreateLaptop function
	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop", err)
		}
		return
	}
	log.Printf("Laptop is created with: %s id", res.Id)
}
