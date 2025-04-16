package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"seotrang.com/rgpc-microservice-product/pkg/grpcclient"
	"seotrang.com/rgpc-microservice-product/services"
)

func main() {
	// Lắng nghe trên cổng 50051
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcclient.InitGRPCConnection(); err != nil {
		log.Fatalf("❌ Failed to connect to gRPC server: %v", err)
	}

	defer func() {
		if grpcclient.UserConn != nil {
			grpcclient.UserConn.Close()
		}
	}()

	// Khởi tạo gRPC server
	s := grpc.NewServer()

	// Đăng ký service từ package services
	services.RegisterServices(s)

	// In ra thông báo và bắt đầu phục vụ
	fmt.Println("🚀 Server is running at :5000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
