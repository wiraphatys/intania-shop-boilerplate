package products

import (
	"context"

	"github.com/wiraphatys/intania-shop-boilerplate/proto/product_pb"
)

type ProductGrpcHandler struct {
	productSvc ProductService
	product_pb.UnimplementedProductServiceServer
}

func NewProductGrpcHandler(productSvc ProductService) *ProductGrpcHandler {
	return &ProductGrpcHandler{
		productSvc: productSvc,
	}
}

func (h *ProductGrpcHandler) GetAllProducts(context.Context, *product_pb.Empty) (*product_pb.GetAllProductsResponse, error) {
	productMock := product_pb.Product{
		Sku:         "INT00001",
		Name:        "Workshop Shirt V1",
		Description: "lorem20 doihgo iopdshgp osdgoi sngodsin sng fgsdh",
		UnitPrice:   290.00,
	}
	return &product_pb.GetAllProductsResponse{
		Products: []*product_pb.Product{
			&productMock,
		},
	}, nil
}

func (h *ProductGrpcHandler) GetProductByID(context.Context, *product_pb.GetProductRequest) (*product_pb.GetProductResponse, error) {
	return nil, nil
}

func (h *ProductGrpcHandler) CreateProduct(context.Context, *product_pb.CreateProductRequest) (*product_pb.Empty, error) {
	return nil, nil
}

func (h *ProductGrpcHandler) UpdateProductByID(context.Context, *product_pb.UpdateProductRequest) (*product_pb.Empty, error) {
	return nil, nil
}

func (h *ProductGrpcHandler) DeleteProductByID(context.Context, *product_pb.DeleteProductRequest) (*product_pb.Empty, error) {
	return nil, nil
}
