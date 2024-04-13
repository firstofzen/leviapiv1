package service

import (
	"levi-apis/apis/response"
	"levi-apis/models/trade"
)

type TradeService struct{}

func (s TradeService) BoothAddProduct(c2 string, id string) error {

}

func (s TradeService) BoothGet(email string) (response.RespBooth, error) {

}

func (s TradeService) BoothUpdateDescription(email string, des string) error {

}

func (s TradeService) OrderAdd(email string, amount float64, category2 string, id string) error {

}

func (s TradeService) OrderGetByEmail(email string) ([]trade.Order, error) {

}

func (s TradeService) CartGet(email string) (trade.Cart, error) {

}

func (s2 TradeService) BoothUpdateBrand(email string, brand string) error {

}
