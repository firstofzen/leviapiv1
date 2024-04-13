package service

import (
	"levi-apis/apis/request"
	"levi-apis/apis/response"
)

type ProductService struct{}

func (s ProductService) GetProductById(c2 string, id string) (response.RespProduct, error) {

}

func (s ProductService) GetProductByCategory2(c2 string) ([]response.RespProduct, error) {

}

func (s ProductService) UpdateSaleProduct(req request.ReqUpdateSaleProduct) error {

}

func (s ProductService) DeleteProduct(c2 string, id string) error {

}

func (s ProductService) GetRating(c2 string, id string) (response.RespProductRating, error) {

}

func (s ProductService) AddRatingUser(req request.ReqAddRatingUser) error {

}

func (s ProductService) RemoveRatingUser(req request.ReqAddRatingUser) error {

}

func (s ProductService) GetProductRanking(c2 string) ([]response.RespProduct, error) {

}

func (s ProductService) AddProductRanking(req request.ReqAddProductRanking) error {

}

func (s ProductService) AddProduct(category2 string, category1 string, description string, sale float64, price float64, title string, avatarUrl string, picturesUrl []string, email string) error {

}
