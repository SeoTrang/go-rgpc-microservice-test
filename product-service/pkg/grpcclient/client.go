package grpcclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"seotrang.com/rgpc-microservice-product/userpb"
)

var UserConn *grpc.ClientConn
var UserClient userpb.UserServiceClient

func InitGRPCConnection() error {
	// Kết nối đến User service (cổng 4000)
	var err error
	UserConn, err = grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// Khởi tạo client cho từng dịch vụ
	UserClient = userpb.NewUserServiceClient(UserConn)
	return nil
}
