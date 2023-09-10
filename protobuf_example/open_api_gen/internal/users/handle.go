package users

import (
	"context"
	pb "github.com/Minsoo-Shin/users/proto/users/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
	"net/http"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (us *UserService) Fetch(ctx context.Context, in *pb.UserFilter) (*pb.UserInfo, error) {
	return &pb.UserInfo{
		Error: nil,
		Users: nil,
	}, nil
}

func (us *UserService) Create(ctx context.Context, in *pb.UserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{
		Error: nil,
		Users: []*pb.User{
			{
				Id:       1,
				Nickname: in.Nickname,
				Email:    in.Email,
				Password: in.Password,
			},
		},
	}, nil
}

// CustomErrorHandler defines the way we want errors to be shown to the articles.
func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	st := status.Convert(err)

	httpStatus := runtime.HTTPStatusFromCode(st.Code())
	w.WriteHeader(httpStatus)

	w.Write([]byte(st.Message()))
}
