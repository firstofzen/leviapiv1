package trade

import (
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/customerror"
	"levi-apis/models/trade"
)

type CartOps struct{}

func (c CartOps) Get(email string) (trade.Cart, error) {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return trade.Cart{}, dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		cartStmt := statement.CartStmt{}
		var rss []trade.Cart
		if err := sess.Query(cartStmt.Get(email), nil).SelectRelease(&rss); err != nil {
			return trade.Cart{}, dbError.QueryError("get cart", err.Error())
		} else {
			if len(rss) > 0 {
				return rss[0], nil
			} else {
				return trade.Cart{}, dbError.RecordNotFound("cart by id")
			}
		}
	}
}

func (c CartOps) UpdateOrderIds(isInsert bool, id string, email string) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		cartStmt := statement.CartStmt{}
		if err := sess.Query(cartStmt.UpdateOrderIds(isInsert, email, id), nil).ExecRelease(); err != nil {
			return dbError.QueryError("update order id", err.Error())
		} else {
			return nil
		}
	}
}

func (c CartOps) Delete(email string) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		cartStmt := statement.CartStmt{}
		if err := sess.Query(cartStmt.Delete(email), nil).ExecRelease(); err != nil {
			return dbError.QueryError("delete cart", err.Error())
		} else {
			return nil
		}
	}
}
