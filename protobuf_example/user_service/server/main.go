package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Minsoo-Shin/protobuf_user/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.CreateUserResponse{
		Name:     in.GetName(),
		Id:       int64(rand.Intn(100000)),
		Email:    in.GetEmail(),
		Phones:   in.GetPhones(),
		CreateDt: timestamppb.Now(),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
