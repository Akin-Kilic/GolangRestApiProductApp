package service

import (
	"errors"
	"go-rest-api/domain"
	"go-rest-api/persistence"
	"go-rest-api/service/model"
)

type IProductSerice interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductSerice {
	return &ProductService{
		productRepository: productRepository,
	}
}
func (productService *ProductService) Add(productCreate model.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}
	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteById(productId)
}

func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}

func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdatePrice(productId, newPrice)
}

func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

func (productService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(storeName)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70 {
		return errors.New("discount cannot be greater than 70")
	}
	return nil
}
