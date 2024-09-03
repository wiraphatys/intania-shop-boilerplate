package products

import (
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
	"go.uber.org/zap"
)

type productServiceImpl struct {
	cfg         config.Config
	logger      *zap.Logger
	productRepo ProductRepository
}

func NewProductService(cfg config.Config, logger *zap.Logger, productRepo ProductRepository) ProductService {
	return &productServiceImpl{
		cfg:         cfg,
		logger:      logger,
		productRepo: productRepo,
	}
}

func (svc *productServiceImpl) GetAllProducts() (*[]Product, error) {
	return nil, nil
}

func (svc *productServiceImpl) GetProductByID(productId string) (*Product, error) {
	return nil, nil
}

func (svc *productServiceImpl) CreateProduct(product *Product) error {
	return nil
}

func (svc *productServiceImpl) UpdateProductByID(productID string, productToUpdate *Product) error {
	return nil
}

func (svc *productServiceImpl) DeleteProductByID(productID string) error {
	return nil
}
