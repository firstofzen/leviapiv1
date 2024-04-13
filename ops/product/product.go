package product

import (
	"levi-apis/constant"
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/customerror"
	"levi-apis/models/product"
)

type ProductOps struct{}

func (o ProductOps) Add(p product.Product) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		productStmt := statement.ProductStmt{}
		if err := sess.Query(productStmt.Add(p.Category2, p.Category1, p.Id.String(), p.Email, p.AvatarUrl, p.PicturesUrl, p.Title, p.Price, p.Sale, p.Inventory), nil).ExecRelease(); err != nil {
			return dbError.QueryError("add product", err.Error())
		} else {
			return nil
		}
	}
}
func (o ProductOps) Get(id string, c2 string) (product.Product, error) {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return product.Product{}, dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		productStmt := statement.ProductStmt{}
		var rss []product.Product
		if err := sess.Query(productStmt.Get(c2, id), nil).SelectRelease(&rss); err != nil {
			return product.Product{}, dbError.QueryError("get product", err.Error())
		} else {
			if len(rss) > 0 {
				return rss[0], nil
			} else {
				return product.Product{}, dbError.RecordNotFound("product by id")
			}
		}
	}
}
func (o ProductOps) UpdateTotalRating(id string, c2 string, totalRating float64) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		pStmt := statement.ProductStmt{}
		if err := sess.Query(pStmt.UpdateTotalRating(c2, id, totalRating), nil).ExecRelease(); err != nil {
			return dbError.QueryError("update total rating", err.Error())
		} else {
			return nil
		}
	}
}
func (o ProductOps) UpdateSale(id string, c2 string, sale float64, ttl int) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		pStmt := statement.ProductStmt{}
		if err := sess.Query(pStmt.UpdateSale(sale, ttl, c2, id), nil).ExecRelease(); err != nil {
			return dbError.QueryError("update sale", err.Error())
		} else {
			return nil
		}
	}
}
func (o ProductOps) Delete(id string, c2 string) error {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbErr.ConnectError(err.Error())
	} else {
		pStmt := statement.ProductStmt{}
		if err := sess.ExecStmt(pStmt.Delete(c2, id)); err != nil {
			return err
		} else {
			return nil
		}
	}
}
func (o ProductOps) GetProductByCategory2(c2 string, page int) ([]product.Product, error) {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return nil, dbErr.ConnectError(err.Error())
	} else {
		var rs []product.Product
		pStmt := statement.ProductStmt{}
		if err := sess.Query(pStmt.GetProductByCategory2(c2, page*constant.PAGE_SIZE), nil).SelectRelease(&rs); err != nil {
			return nil, err
		} else {
			return rs, nil
		}
	}
}
func (o ProductOps) GetCountProductByCategory2(c2 string) (int, error) {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return 0, dbErr.ConnectError(err.Error())
	} else {
		var rs int
		pStmt := statement.ProductStmt{}
		if err := sess.Query(pStmt.GetCountProductByCategory2(c2), nil).SelectRelease(&rs); err != nil {
			return 0, dbErr.QueryError("get count product by category2", err.Error())
		} else {
			return rs, nil
		}
	}
}
