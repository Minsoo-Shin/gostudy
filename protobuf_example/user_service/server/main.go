package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Minsoo-Shin/protobuf_user/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strings"
)

var (
	port = flag.Int("port", 3000, "The server port")
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

func allHandler(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request", fmt.Sprintf("%+v", r))
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
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
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions:   protojson.MarshalOptions{},
				UnmarshalOptions: protojson.UnmarshalOptions{},
			}),
	)

	err = pb.RegisterUserHandlerFromEndpoint(context.Background(), gwmux, lis.Addr().String(), []grpc.DialOption{
		grpc.WithInsecure(),
	})
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(lis.Addr().String()+"1", allHandler(s, gwmux))
	if err != nil {
		panic(err)
	}
}
