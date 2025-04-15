package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"seotrang.com/rgpc-microservice/pkg/grpcclient"
	"seotrang.com/rgpc-microservice/routes"
)

func main() {
	// Khởi tạo kết nối gRPC cho cả hai dịch vụ
	if err := grpcclient.InitGRPCConnection(); err != nil {
		log.Fatalf("❌ Failed to connect to gRPC server: %v", err)
	}
	defer func() {
		if grpcclient.UserConn != nil {
			grpcclient.UserConn.Close()
		}
		if grpcclient.ProductConn != nil {
			grpcclient.ProductConn.Close()
		}
	}()

	router := gin.Default()
	routes.SetupRoutes(router)
	log.Println("🚀 Server running on http://localhost:3000")
	router.Run(":3000")
}
