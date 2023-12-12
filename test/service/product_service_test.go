package service

import (
	"go-rest-api/domain"
	"go-rest-api/service"
	"go-rest-api/service/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductSerice

func TestMain(m *testing.M) {

	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "ABC TECH",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		actualProduct := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProduct))
	})
}

func Test_WhenNoValidationErrorOccurred_ShoukdAddProduct(t *testing.T) {
	t.Run("WhenNoValidationErrorOccurred_ShoukdAddProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		}, actualProducts[len(actualProducts)-1])
	})
}

func Test_WhenDiscountOccurred_ShoukdAddProduct(t *testing.T) {
	t.Run("WhenNoValidationErrorOccurred_ShoukdAddProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		}, actualProducts[len(actualProducts)-1])
	})
}
