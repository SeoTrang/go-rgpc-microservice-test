package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"seotrang.com/rgpc-microservice-product/helpers"
	"seotrang.com/rgpc-microservice-product/internal/data"
	"seotrang.com/rgpc-microservice-product/pkg/grpcclient"
	"seotrang.com/rgpc-microservice-product/productpb"
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

	productRes
	for _, product := range data.MockProducts {
		if product.Id == id {

			return &productpb.Product{
				Id:    product.Id,
				Name:  product.Name,
				Price: product.Price,
			}, nil
		}
	}

	ctx, cancel := helpers.NewCtx()
	defer cancel()

	res, err := grpcclient.UserClient.GetUser(productRes.user_id)

	return nil, status.Errorf(codes.NotFound, "Product with ID %d not found", id)
}

// GetAllProduct trả về danh sách tất cả sản phẩm
func (s *ProductService) GetAllProduct(ctx context.Context, req *productpb.GetAllProductRequest) (*productpb.GetAllProductResponse, error) {
	var productList []*productpb.Product

	for _, product := range data.MockProducts {
		productList = append(productList, &productpb.Product{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return &productpb.GetAllProductResponse{
		Product: productList,
	}, nil
}
