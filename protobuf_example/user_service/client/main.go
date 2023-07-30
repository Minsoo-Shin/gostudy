package main

import (
	"context"
	"flag"
	pb "github.com/Minsoo-Shin/protobuf_user/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultEmail       = "default@gmail.com"
	defaultPassword    = "1234"
	defaultName        = "minsoo"
	defaultPhoneNumber = "010-0000-0000"
	defaultPhoneType   = 0
)

var (
	addr        = flag.String("addr", "localhost:50051", "the address to connect to")
	name        = flag.String("name", defaultName, "Name")
	email       = flag.String("email", defaultEmail, "Email")
	password    = flag.String("password", defaultPassword, "Password")
	phoneNumber = flag.String("phone", defaultPhoneNumber, "PhoneNumber")
	phoneType   = flag.Int("phoneType", defaultPhoneType, "PhoneType")
)

func main() {
	flag.Parse()
	log.Printf("addr %v", *addr)
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.CreateUserRequest{
		Email:    *email,
		Password: *password,
		Name:     *name,
		Phones: []*pb.PhoneNumber{
			{
				Number: *phoneNumber,
				Type:   pb.PhoneType(*phoneType),
			},
		},
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Server said: User %v Created", r.GetId())
}
