package data

import "seotrang.com/rgpc-microservice-product/models"

// MockUsers chứa dữ liệu mẫu user
var MockProducts = []models.Product{
	{Id: 1, Name: "iPhone 15", UserId: 1, Price: 999.99},
	{Id: 2, Name: "Samsung Galaxy S24", UserId: 2, Price: 899.99},
	{Id: 3, Name: "Xiaomi Mi 13", UserId: 1, Price: 499.99},
	{Id: 4, Name: "MacBook Pro M3", UserId: 2, Price: 2199.00},
	{Id: 5, Name: "Dell XPS 13", UserId: 3, Price: 1299.49},
}
