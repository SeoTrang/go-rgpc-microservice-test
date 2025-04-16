package grpcclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"seotrang.com/rgpc-microservice/orderpb"
	"seotrang.com/rgpc-microservice/productpb"
	"seotrang.com/rgpc-microservice/userpb"
)

var UserConn *grpc.ClientConn
var ProductConn *grpc.ClientConn
var OrderConn *grpc.ClientConn
var UserClient userpb.UserServiceClient
var ProductClient productpb.ProductServiceClient
var OrderClient orderpb.OrderServiceClient

func InitGRPCConnection() error {
	// Kết nối đến User service (cổng 4000)
	var err error
	UserConn, err = grpc.Dial("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// Kết nối đến Product service (cổng 5000)
	ProductConn, err = grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	OrderConn, err = grpc.Dial("localhost:6000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// Khởi tạo client cho từng dịch vụ
	UserClient = userpb.NewUserServiceClient(UserConn)
	ProductClient = productpb.NewProductServiceClient(ProductConn)
	OrderClient = orderpb.NewOrderServiceClient(OrderConn)

	return nil
}
