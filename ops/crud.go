package ops

import (
	"levi-apis/ops/product"
	"levi-apis/ops/trade"
	"levi-apis/ops/user"
)

type CrudStdOps struct {
}

func (o CrudStdOps) NewAccountOps() *user.AccountOps {
	return &user.AccountOps{}
}
func (o CrudStdOps) NewUserDetailOps() *user.DetailOps {
	return &user.DetailOps{}
}
func (o CrudStdOps) NewProductOps() *product.ProductOps {
	return &product.ProductOps{}
}
func (o CrudStdOps) NewCartOps() *trade.CartOps {
	return &trade.CartOps{}
}
