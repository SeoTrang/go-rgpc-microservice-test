package services

import (
	"google.golang.org/grpc"
	"seotrang.com/rgpc-microservice-user/userpb"
)

// RegisterServices đăng ký tất cả các service với gRPC server
func RegisterServices(s *grpc.Server) {
	// Đăng ký UserService
	userpb.RegisterUserServiceServer(s, NewUserService())

	// Đăng ký các service khác ở đây, ví dụ:
	// otherpb.RegisterOtherServiceServer(s, NewOtherService())
}
