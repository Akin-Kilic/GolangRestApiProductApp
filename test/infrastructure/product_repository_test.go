package infrastructure

import (
	"context"
	"go-rest-api/common/postgresql"
	"go-rest-api/domain"
	"go-rest-api/persistence"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		UserName:              "postgres",
		Password:              "postgres",
		DbName:                "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = persistence.NewProductRepository(dbPool)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}
func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProduct := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProduct))
		assert.Equal(t, expectedProducts, actualProduct)
	})
	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProduct := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProduct))
		assert.Equal(t, expectedProducts, actualProduct)
	})
	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "Kupa",
			Price:    100.0,
			Discount: 0.0,
			Store:    "Kırtasiye Merkezi",
		},
	}

	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    100.0,
		Discount: 0.0,
		Store:    "Kırtasiye Merkezi",
	}
	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetbyId(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		_, err := productRepository.GetById(5)
		assert.Equal(t, domain.Product{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		}, actualProduct)
		assert.Equal(t, "product not found by id", err.Error())

	})
	clear(ctx, dbPool)
}

func TestDeleteById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("DeleteById", func(t *testing.T) {
		productRepository.DeleteById(1)
		_, err := productRepository.GetById(1)
		assert.Equal(t, "product not found with id", err.Error())
	})
	clear(ctx, dbPool)
}

func TestUpdatePrice(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("UpdatePrice", func(t *testing.T) {
		productBeforeUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, float32(3000.0), productBeforeUpdate.Price)

		productRepository.UpdatePrice(1, 4000.0)
		productAfterUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, float32(4000.0), productAfterUpdate.Price)
	})
	clear(ctx, dbPool)
}
