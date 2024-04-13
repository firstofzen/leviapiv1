package trade

import (
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/customerror"
	"levi-apis/models/trade"
	"strconv"
)

type BoothOps struct{}

func (b BoothOps) Get(email string) (trade.Booth, error) {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return trade.Booth{}, dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		bStmt := statement.BoothStmt{}
		var rss []trade.Booth
		if err := sess.Query(bStmt.Get(email), nil).SelectRelease(&rss); err != nil {
			return trade.Booth{}, dbError.QueryError("get booth", err.Error())
		} else {
			if len(rss) > 0 {
				return rss[0], nil
			} else {
				return trade.Booth{}, dbError.RecordNotFound("get booth")
			}
		}
	}
}
func (b BoothOps) UpdateDescription(email string, des string) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		bStmt := statement.BoothStmt{}
		if err := sess.ExecStmt(bStmt.UpdateAField(email, "description", des)); err != nil {
			return dbError.QueryError("update description", err.Error())
		} else {
			return nil
		}
	}
}
func (b BoothOps) UpdateAuthentic(email string, isAuthentic bool) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		bStmt := statement.BoothStmt{}
		if err := sess.ExecStmt(bStmt.UpdateAField(email, "authentic", strconv.FormatBool(isAuthentic))); err != nil {
			return dbError.QueryError("update authentic", err.Error())
		} else {
			return nil
		}
	}
}
func (b BoothOps) UpdateMainCategory(c string, email string) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		bStmt := statement.BoothStmt{}
		if err := sess.ExecStmt(bStmt.UpdateAField(email, "main_category", c)); err != nil {
			return dbError.QueryError("update main category", err.Error())
		} else {
			return nil
		}
	}
}
func (b BoothOps) Delete(email string) error {
	var dbError customerror.DBError
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		bStmt := statement.BoothStmt{}
		if err := sess.ExecStmt(bStmt.Delete(email)); err != nil {
			return dbError.QueryError("delete booth", err.Error())
		} else {
			return nil
		}
	}
}
