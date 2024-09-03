package products

import "gorm.io/gorm"

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

func (repo *productRepositoryImpl) FindAllProducts() (*[]Product, error) {
	return nil, nil
}
