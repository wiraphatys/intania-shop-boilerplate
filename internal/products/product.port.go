package products

type ProductService interface {
	GetAllProducts() (*[]Product, error)
	GetProductByID(productID string) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProductByID(productID string, productToUpdate *Product) error
	DeleteProductByID(productID string) error
}

type ProductRepository interface {
	FindAllProducts() (*[]Product, error)
}
