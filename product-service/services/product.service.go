package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"seotrang.com/rgpc-microservice-product/helpers"
	"seotrang.com/rgpc-microservice-product/internal/data"
	"seotrang.com/rgpc-microservice-product/models"
	"seotrang.com/rgpc-microservice-product/pkg/grpcclient"
	"seotrang.com/rgpc-microservice-product/productpb"
	"seotrang.com/rgpc-microservice-product/userpb"
)

// ProductService implements the ProductService gRPC server
type ProductService struct {
	productpb.UnimplementedProductServiceServer
}

// NewProductService returns a new instance of ProductService
func NewProductService() *ProductService {
	return &ProductService{}
}

// GetProduct trả về thông tin một sản phẩm theo ID
func (s *ProductService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.Product, error) {
	id := req.GetId()

	// Tìm sản phẩm theo ID
	var productRes *models.Product
	for _, product := range data.MockProducts {
		if product.Id == id {
			productRes = &models.Product{
				Id:     product.Id,
				Name:   product.Name,
				Price:  product.Price,
				UserId: product.UserId,
			}
			break
		}
	}

	// Nếu không tìm thấy sản phẩm
	if productRes == nil {
		return nil, status.Errorf(codes.NotFound, "Product with ID %d not found", id)
	}

	// Gọi sang service User để lấy thông tin user
	userCtx, cancel := helpers.NewCtx()
	defer cancel()

	res, err := grpcclient.UserClient.GetUser(userCtx, &userpb.GetUserRequest{Id: productRes.UserId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not fetch user with id %d: %v", productRes.UserId, err)
	}

	// Gán user vào productRes
	productRes.User = &models.User{
		ID:   res.Id,
		Name: res.Name,
		Age:  res.Age,
	}

	// Trả về protobuf Product (chuyển đổi từ model sang pb)
	return &productpb.Product{
		Id:    productRes.Id,
		Name:  productRes.Name,
		Price: productRes.Price,
		User: &productpb.ProductUser{
			Id:   res.Id,
			Name: res.Name,
			Age:  res.Age,
		},
	}, nil
}

// GetAllProduct trả về danh sách tất cả sản phẩm
func (s *ProductService) GetAllProduct(ctx context.Context, req *productpb.GetAllProductRequest) (*productpb.GetAllProductResponse, error) {
	var productList []*productpb.Product

	for _, product := range data.MockProducts {
		// Gọi sang service User để lấy thông tin user
		userCtx, cancel := helpers.NewCtx()
		defer cancel()

		res, err := grpcclient.UserClient.GetUser(userCtx, &userpb.GetUserRequest{Id: product.UserId})
		if err != nil {
			// Nếu không lấy được user, có thể bỏ qua user hoặc return lỗi tùy nhu cầu
			continue
		}

		productList = append(productList, &productpb.Product{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
			User: &productpb.ProductUser{
				Id:   res.Id,
				Name: res.Name,
				Age:  res.Age,
			},
		})
	}

	return &productpb.GetAllProductResponse{
		Product: productList,
	}, nil
}
