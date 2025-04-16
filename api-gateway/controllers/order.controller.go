package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"seotrang.com/rgpc-microservice/helpers"
	"seotrang.com/rgpc-microservice/models"
	"seotrang.com/rgpc-microservice/orderpb"
	"seotrang.com/rgpc-microservice/pkg/grpcclient"
)

func GetOrder(c *gin.Context) {
	ctx, cancel := helpers.NewCtx()
	defer cancel()

	// Lấy ID từ route param
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	// Gọi gRPC đến OrderService
	res, err := grpcclient.OrderClient.GetOrder(ctx, &orderpb.OrderRequest{Id: int32(id)})
	if err != nil {
		log.Println("❌ could not found order with id:", id, err)
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	order := models.Order{
		Id:     res.Id,
		Code:   res.Code,
		UserId: res.UserId,
		User: &models.User{
			ID:   res.User.Id,
			Name: res.User.Name,
			Age:  res.User.Age,
		},
	}

	c.JSON(200, order)
}

func GetAllOrders(c *gin.Context) {
	ctx, cancel := helpers.NewCtx()
	defer cancel()

	// Gọi gRPC đến OrderService
	res, err := grpcclient.OrderClient.GetAllOrders(ctx, &orderpb.Empty{})
	if err != nil {
		log.Println("❌ could not fetch orders:", err)
		c.JSON(500, gin.H{"error": "Could not fetch orders"})
		return
	}

	var orders []models.Order
	for _, o := range res.Orders {
		order := models.Order{
			Id:     o.Id,
			Code:   o.Code,
			UserId: o.UserId,
			User: &models.User{
				ID:   o.User.Id,
				Name: o.User.Name,
				Age:  o.User.Age,
			},
		}
		orders = append(orders, order)
	}

	c.JSON(200, orders)
}
