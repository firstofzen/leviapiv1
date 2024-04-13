package product

import (
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/customerror"
	"levi-apis/models/product"
)

type RatingOps struct {
}

func (o RatingOps) AddRating(c2 string, id string) error {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbErr.ConnectError(err.Error())
	} else {
		rStmt := statement.RatingStmt{}
		if err := sess.ExecStmt(rStmt.Add(c2, id)); err != nil {
			return dbErr.QueryError("add rating", err.Error())
		} else {
			return nil
		}
	}
}
func (o RatingOps) GetRating(c2 string, id string) (product.Rating, error) {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return product.Rating{}, dbErr.ConnectError(err.Error())
	} else {
		rStmt := statement.RatingStmt{}
		var rs []product.Rating
		if err := sess.Query(rStmt.Get(c2, id), nil).SelectRelease(&rs); err != nil {
			return product.Rating{}, dbErr.QueryError("get rating", err.Error())
		} else {
			if len(rs) > 0 {
				return rs[0], nil
			} else {
				return product.Rating{}, dbErr.RecordNotFound("rating")
			}
		}
	}
}
