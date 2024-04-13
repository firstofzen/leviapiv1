package user

import (
	"levi-apis/db"
	"levi-apis/db/statement"
	"levi-apis/models/user"
)

type DetailOps struct {
}

func (u DetailOps) UpdateAField(name string, value string, email string) error {
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		detailStmt := statement.UserDetailStmt{}
		if err := sess.ExecStmt(detailStmt.UpdateField(name, value, email)); err != nil {
			return dbError.QueryError("update field", err.Error())
		} else {
			return nil
		}
	}
}
func (u DetailOps) Get(email string) (user.Detail, error) {
	if sess, err := db.Connect(); err != nil {
		return user.Detail{}, dbError.ConnectError(err.Error())
	} else {
		defer sess.Close()
		detailStmt := statement.UserDetailStmt{}
		var rss []user.Detail
		if err := sess.Query(detailStmt.Get(email), nil).SelectRelease(&rss); err != nil {
			return user.Detail{}, dbError.QueryError("get user detail", err.Error())
		} else {
			if len(rss) > 0 {
				return rss[0], nil
			} else {
				return user.Detail{}, dbError.RecordNotFound("get user detail by id")
			}
		}
	}
}
func (u DetailOps) Delete(email string) error {
	if sess, err := db.Connect(); err != nil {
		return dbError.ConnectError(err.Error())
	} else {
		detailStmt := statement.UserDetailStmt{}
		if err := sess.ExecStmt(detailStmt.Delete(email)); err != nil {
			return dbError.QueryError("delete user detail", err.Error())
		} else {
			return nil
		}
	}
}
