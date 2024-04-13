package trade

import (
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/customerror"
	"levi-apis/models/trade"
)

type OrderOps struct{}

func (o OrderOps) Add(email string, id string, pc2 string, pid string, totalAmount float64) error {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbErr.ConnectError(err.Error())
	} else {
		oStmt := statement.OrderStmt{}
		if err := sess.ExecStmt(oStmt.Add(email, id, pc2, pid, totalAmount)); err != nil {
			return dbErr.QueryError("add order", err.Error())
		} else {
			return nil
		}
	}
}
func (o OrderOps) GetByEmail(email string) ([]trade.Order, error) {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return nil, dbErr.ConnectError(err.Error())
	} else {
		oStmt := statement.OrderStmt{}
		var rs []trade.Order
		if err := sess.Query(oStmt.GetByEmail(email), nil).SelectRelease(&rs); err != nil {
			return nil, dbErr.QueryError("get order by email", err.Error())
		} else {
			return rs, nil
		}
	}
}
func (o OrderOps) UpdateNextStatus(email string, id string, status string) error {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbErr.ConnectError(err.Error())
	} else {
		var oStmt statement.OrderStmt
		if err := sess.ExecStmt(oStmt.UpdateStatus(status, email, id)); err != nil {
			return dbErr.QueryError("update status", err.Error())
		} else {
			return nil
		}
	}
}
func (o OrderOps) Delete(email string, id string) error {
	var dbErr customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbErr.ConnectError(err.Error())
	} else {
		var oStmt statement.OrderStmt
		if err := sess.ExecStmt(oStmt.Delete(email, id)); err != nil {
			return dbErr.QueryError("delete order", err.Error())
		} else {
			return nil
		}
	}
}
