package statement

import (
	"fmt"
)

type AccountStmt struct {
}

func (a AccountStmt) Add(avatar string, name string, email string) string {
	return fmt.Sprintf("insert into user.account(email, name, avatar_url) values ('%s', '%s', '%s')", email, name, avatar)
}
func (a AccountStmt) Get(email string) string {
	return fmt.Sprintf("select * from user.account where email='%s'", email)
}
func (a AccountStmt) Delete(email string) string {
	return fmt.Sprintf("delete from user.account where email='%s'", email)
}
func (a AccountStmt) UpdateField(fieldName string, fieldValue string, email string) string {
	return fmt.Sprintf("update user.account set %s = %s where email = '%s'", fieldName, fieldValue, email)
}
func (a AccountStmt) CheckExist(email string) string {
	return fmt.Sprintf("select email from user.account where email='%s'", email)
}

type UserDetailStmt struct {
}

func (u UserDetailStmt) Add(email string, phone string, address string) string {
	return fmt.Sprintf("insert into user.detail(email, phone, address, is_violate) values ('%s', '%s', '%s', false)", email, phone, address)
}
func (u UserDetailStmt) Get(email string) string {
	return fmt.Sprintf("select * from user.detail where email='%s'", email)
}
func (u UserDetailStmt) UpdateField(fieldName string, fieldValue string, email string) string {
	return fmt.Sprintf("update user.detail set %s = %s where email='%s'", fieldName, fieldValue, email)
}
func (u UserDetailStmt) Delete(email string) string {
	return fmt.Sprintf("delete from user.detail where email='%s'", email)
}
