package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "example.com/laptop_store/proto"
	"example.com/laptop_store/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func seedUser(store service.UserStore) error {
	err := createUser(store, "Milad", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(store, "Jafar", "secret", "user")
}

func createUser(userStore service.UserStore, username string, password string, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return fmt.Errorf("Cannot create a User: %v", err)
	}
	return userStore.Save(user)
}

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/laptostore.grpc.LaptopService/"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
	}
}
func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	userStore := service.NewInMemoryUserStore()
	jwrManger := service.NewJWTManager(secretKey, tokenDuration)
	err := seedUser(userStore)
	if err != nil {
		log.Fatal("Cannot seed user")
	}
	authServer := service.NewAuthServer(userStore, jwrManger)
	interceptor := service.NewAuthInterceptor(jwrManger, accessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}
