package persistence

import (
	"context"
	"errors"
	"fmt"
	"go-rest-api/domain"
	"go-rest-api/persistence/common"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "select * from products")
	if err != nil {
		log.Errorf("Error while getting all products: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}
func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()
	getProductsByStoreNameSql := `select * from products where store =$1`
	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)
	if err != nil {
		log.Errorf("Error while getting all products: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRows)
}
func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `Insert into products (name,price,discount,store) VALUES ($1,$2,$3,$4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Failed to add new product", err)
		return err
	}
	log.Info(fmt.Printf("Product added with %v", addNewProduct))
	return nil
}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getByIdSql := `select * from products where id =$1`
	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)
	return extractProductFromRow(queryRow)

}
func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()
	_, getErr := productRepository.GetById(productId)
	if getErr != nil {
		return errors.New("product not found")
	}
	deleteSql := `delete from products where id =$1`
	_, err := productRepository.dbPool.Exec(ctx, deleteSql, productId)
	if err != nil {
		x := fmt.Sprintf("error while deleting product with id %v", productId)
		return errors.New(x)
	}
	log.Info("Deleted item successfully!")
	return nil

}

func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	updateSql := "update products set price = $1 where id = $2"
	_, err := productRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)
	if err != nil {
		return errors.New("error while updating product with id")
	}
	log.Info("Updated product price successfully")
	return nil

}
func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products []domain.Product
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}

	return products
}
func extractProductFromRow(productRow pgx.Row) (domain.Product, error) {
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := productRow.Scan(&id, &name, &price, &discount, &store)

	if scanErr != nil && scanErr.Error() == common.NOT_FOUND {
		return domain.Product{}, fmt.Errorf("product not found by id")
	}
	if scanErr != nil {
		return domain.Product{}, fmt.Errorf("error while getting product by id")
	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}
