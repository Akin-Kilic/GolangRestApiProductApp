package service

import (
	"errors"
	"go-rest-api/domain"
	"go-rest-api/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}
func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	var storeProducts []domain.Product

	for _, product := range fakeRepository.products {
		if product.Store == storeName {
			storeProducts = append(storeProducts, product)
		}
	}

	return storeProducts
}
func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	for _, product := range fakeRepository.products {
		if product.Id == productId {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}
func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {

	var foundIndex = -1

	for i, product := range fakeRepository.products {
		if product.Id == productId {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return errors.New("product not found")
	}

	fakeRepository.products = append(fakeRepository.products[:foundIndex], fakeRepository.products[foundIndex+1:]...)

	return nil

}

func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	for index, product := range fakeRepository.products {
		if product.Id == productId {
			fakeRepository.products[index].Price = newPrice
			return nil
		}
	}
	return errors.New("product not found")

}
