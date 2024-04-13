package user

import (
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/customerror"
	"levi-apis/models/user"
)

type AccountOps struct {
}

var dbError customerror.DBError

func (o AccountOps) Init(avatarUrl string, name string, email string) error {
	accStmt := statement.AccountStmt{}
	cartStmt := statement.CartStmt{}
	detailStmt := statement.UserDetailStmt{}
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		if err := sess.ExecStmt(statement.BatchStmt(
			accStmt.Add(avatarUrl, name, email),
			cartStmt.Add(email),
			detailStmt.Add(email, "", ""))); err != nil {
			return dbError.QueryError("init account", err.Error())
		} else {
			return nil
		}
	}
}

func (o AccountOps) Get(email string) (user.Account, error) {
	if sess, err := db.Connect(); err != nil {
		return user.Account{}, dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		accStmt := statement.AccountStmt{}
		var rs []user.Account
		if err := sess.Query(accStmt.Get(email), nil).SelectRelease(&rs); err != nil {
			return user.Account{}, dbError.QueryError("get account by email", err.Error())
		} else {
			if len(rs) == 0 {
				return user.Account{}, dbError.RecordNotFound("get account by email")
			} else {
				return rs[0], nil
			}
		}
	}
}
func (o AccountOps) UpdateAField(fieldName string, fieldValue string, email string) error {
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		accStmt := statement.AccountStmt{}
		if err := sess.ExecStmt(accStmt.UpdateField(fieldName, "'"+fieldValue+"'", email)); err != nil {
			return dbError.QueryError("update field", err.Error())
		} else {
			return nil
		}
	}
}
func (o AccountOps) Delete(email string) error {
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		accStmt := statement.AccountStmt{}
		if err := sess.ExecStmt(accStmt.Delete(email)); err != nil {
			return dbError.QueryError("delete", err.Error())
		} else {
			return nil
		}
	}
}

func (o AccountOps) CheckExists(email string) (bool, error) {
	if sess, err := db.Connect(); err != nil {
		return false, dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		accStmt := statement.AccountStmt{}
		var rs []string
		if err := sess.Query(accStmt.CheckExist(email), nil).SelectRelease(&rs); err != nil {
			return false, dbError.QueryError("check account exist", err.Error())
		} else {
			if len(rs) > 0 {
				return true, nil
			} else {
				return false, nil
			}
		}
	}
}
