package grpc

import (
	"context"
	"errors"

	"github.com/atefeh-syf/grpc-clean/internal/models"
	interfaces "github.com/atefeh-syf/grpc-clean/pkg/v1"
	pb "github.com/atefeh-syf/grpc-clean/proto"

	//userRepo "github.com/atefeh-syf/grpc-clean/pkg/v1/repository"
	"google.golang.org/grpc"
)

type UserServStruct struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
  }
 
 func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface){
   userGrpc := &UserServStruct{useCase: usecase}
   pb.RegisterUserServiceServer(grpcServer, userGrpc)
 }

func (srv *UserServStruct) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.UserProfileResponse{}, errors.New("please provide all fields")
	}

	user, err := srv.useCase.Create(data)

	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

func (srv *UserServStruct) Get(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
	  return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}
	user, err := srv.useCase.Get(id)
	if err != nil {
	  return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}
  

func (srv *UserServStruct) Update(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
	  return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}
	user, err := srv.useCase.Get(id)
	if err != nil {
	  return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

func (srv *UserServStruct) Delete(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
	  return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}
	user, err := srv.useCase.Get(id)
	if err != nil {
	  return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}


func (srv *UserServStruct) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}


func (srv *UserServStruct) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Id: string(user.ID), Name: user.Name, Email: user.Email}
}