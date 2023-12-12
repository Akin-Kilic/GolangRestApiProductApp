package response

import "go-rest-api/domain"

type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"`
}

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

func ToResponseList(products []domain.Product) []ProductResponse {
	var productsResponseList = []ProductResponse{}
	for _, product := range products {
		productsResponseList = append(productsResponseList, ToResponse(product))
	}
	return productsResponseList
}
