package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"seotrang.com/rgpc-microservice-user/internal/data"
	"seotrang.com/rgpc-microservice-user/userpb"
)

// Define Server in the services package
type UserService struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUser lấy thông tin user theo ID
func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	id := req.GetId()

	for _, user := range data.MockUsers {
		if user.ID == id {
			return &userpb.User{
				Id:   user.ID,
				Name: user.Name,
				Age:  user.Age,
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "User with ID %d not found", id)
}

// GetAllUsers lấy danh sách tất cả user
func (s *UserService) GetAllUsers(ctx context.Context, req *userpb.GetAllUserRequest) (*userpb.GetAllUserResponse, error) {
	var userList []*userpb.User

	for _, user := range data.MockUsers {
		userList = append(userList, &userpb.User{
			Id:   user.ID,
			Name: user.Name,
			Age:  user.Age,
		})
	}

	return &userpb.GetAllUserResponse{
		Users: userList,
	}, nil
}
