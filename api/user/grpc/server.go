package grpc

import (
	"context"

	"github.com/alizoubair/go-grpc-server/api/api_struct/form"
	"github.com/alizoubair/go-grpc-server/api/user"
	"github.com/alizoubair/go-grpc-server/api/user/proto"
	ts "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rs/xid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedUserServiceServer
	svc user.UserService
}

func NewUserServerGrpc(s *grpc.Server, svc user.UserService) *server {
	usrSvr := &server{
		UnimplementedUserServiceServer: proto.UnimplementedUserServiceServer{},
		svc:                            svc,
	}
	proto.RegisterUserServiceServer(s, usrSvr)
	reflection.Register(s)
	return usrSvr
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	userReq := &form.UserForm{
		ID:      xid.New().String(),
		Name:    req.GetName(),
		Email:   req.GetEmail(),
		Address: req.GetAddress(),
		Phone:   req.GetPhone(),
	}

	user, err := s.svc.CreateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}

	createdAt := &ts.Timestamp{
		Seconds: user.CreatedAt.Unix(),
	}

	updatedAt := &ts.Timestamp{
		Seconds: user.UpdatedAt.Unix(),
	}

	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Address:   user.Address,
			Phone:     user.Phone,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}, nil
}
