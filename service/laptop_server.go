package service

import(
	"context"
	"log"
	"errors"

	"example.com/laptop_store/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type LaptopServer struct{
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}
func NewLaptopServer(store LaptopStore) *LaptopServer{
	return &LaptopServer{Store: 	store}
}

func (server LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error){
	laptop := req.GetLaptop()
	log.Printf("Receive create_laptop request")

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		
		if err != nil{
			return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not a valid UUID: %v", err)
		}
	}else{
		id, err := uuid.NewRandom()
		if err != nil{
			return nil, status.Errorf(codes.Internal, "Cannot generate a valid UUID: %v", err)
		}
		laptop.Id = id.String()
	}

	// Check if the request timeout is exceeded
	if ctx.Err() == context.DeadlineExceeded{
		log.Printf("Deadline is exceeded")

		return nil, status.Errorf(codes.DeadlineExceeded, "Deadline is exceeded")
	}

	// Check if the request is cancelled by client 
	if ctx.Err() == context.Canceled{
		log.Printf("Request is canceled")

		return nil, status.Errorf(codes.Canceled, "Request is canceled")
	}
	// Save data to the db. 
	// In this project the main focus is on the gRPC so we use in-memory storage. 
	err := server.Store.Save(laptop)
	if err != nil{
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists){
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "Cannot save laptop data to store, %v", err)
	}

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}