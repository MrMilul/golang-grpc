package main

import (
	"log"
	"flag"
	"net"
	"fmt"

	"google.golang.org/grpc"
	"example.com/laptop_store/service"
	"example.com/laptop_store/proto"
)

func main(){
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	grpcServer := grpc.NewServer()

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil{
		log.Fatal("Cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil{
			log.Fatal("Cannot start server: ", err)
		}
	

}