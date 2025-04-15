package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"seotrang.com/rgpc-microservice/helpers"
	"seotrang.com/rgpc-microservice/models"
	"seotrang.com/rgpc-microservice/pkg/grpcclient"
	"seotrang.com/rgpc-microservice/userpb"
)

func GetUser(c *gin.Context) {
	ctx, cancel := helpers.NewCtx()
	println("vo")
	defer cancel()
	res, err := grpcclient.UserClient.GetUser(ctx, &userpb.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not found user with id 1")
	}

	user := models.User{
		ID:   res.Id,
		Name: res.Name,
		Age:  res.Age,
	}

	c.JSON(200, user)

}

func GetUsers(c *gin.Context) {
	// Tạo context và hủy khi hoàn thành
	ctx, cancel := helpers.NewCtx()
	defer cancel()

	// Gọi GetAllUsers từ gRPC
	res, err := grpcclient.UserClient.GetAllUsers(ctx, &userpb.GetAllUserRequest{})
	if err != nil {
		log.Fatalf("❌ could not found users: %v", err)
		// Có thể trả về 500 lỗi cho client để họ biết có vấn đề
		c.JSON(500, gin.H{"error": "could not fetch users"})
		return
	}

	// Chuyển đổi response thành slice của models.User
	users := []models.User{}
	for _, u := range res.Users {
		users = append(users, models.User{
			ID:   u.Id,
			Name: u.Name,
			Age:  u.Age,
		})
	}

	// Trả về danh sách users dưới dạng JSON
	c.JSON(200, users)
}
