package service

import(
	"context"


	"example.com/laptop_store/proto"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type AuthServer struct{
	store UserStore
	jwtManager *JWTManager
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(store UserStore, jwtManager *JWTManager)*AuthServer{
	return &AuthServer{
		store: store,
		jwtManager: jwtManager,}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error){
	user, err := server.store.Find(req.GetUsername())
	if err != nil{
		return nil, status.Errorf(codes.Internal, "Cannot find user %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()){
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")

	}

	token, err := server.jwtManager.TokenGenerator(user)
	if err != nil{
		return nil, status.Errorf(codes.Internal, "Cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}