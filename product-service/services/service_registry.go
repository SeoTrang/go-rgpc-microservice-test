package services

import (
	"google.golang.org/grpc"
	"seotrang.com/rgpc-microservice-product/productpb"
)

// RegisterServices đăng ký tất cả các service với gRPC server
func RegisterServices(s *grpc.Server) {
	// Đăng ký UserService
	productpb.RegisterProductServiceServer(s, NewProductService())
	
	// Đăng ký các service khác ở đây, ví dụ:
	// otherpb.RegisterOtherServiceServer(s, NewOtherService())
}
