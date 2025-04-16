package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"seotrang.com/rgpc-microservice/helpers"
	"seotrang.com/rgpc-microservice/models"
	"seotrang.com/rgpc-microservice/pkg/grpcclient"
	"seotrang.com/rgpc-microservice/productpb"
)

func GetProduct(c *gin.Context) {
	ctx, cancel := helpers.NewCtx()
	defer cancel()
	res, err := grpcclient.ProductClient.GetProduct(ctx, &productpb.GetProductRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not found product with id 1")
	}

	product := models.Product{
		Id:     res.Id,
		Name:   res.Name,
		Price:  res.Price,
		UserId: res.UserId,
		User: &models.User{
			ID:   res.User.Id,
			Name: res.User.Name,
			Age:  res.User.Age,
		},
	}

	c.JSON(200, product)
}

func GetAllProducts(c *gin.Context) {
	ctx, cancel := helpers.NewCtx()
	defer cancel()
	res, err := grpcclient.ProductClient.GetAllProduct(ctx, &productpb.GetAllProductRequest{})
	if err != nil {
		log.Println(err)
		log.Fatalf("could not found products")
	}

	var products []models.Product
	for _, p := range res.Product {
		product := models.Product{
			Id:     p.Id,
			Name:   p.Name,
			Price:  p.Price,
			UserId: p.UserId,
			User: &models.User{
				ID:   p.User.Id,
				Name: p.User.Name,
				Age:  p.User.Age,
			},
		}
		products = append(products, product)
	}

	c.JSON(200, products)
}
