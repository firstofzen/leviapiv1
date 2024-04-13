package service

import (
	user2 "levi-apis/models/user"
	"levi-apis/ops/user"
)

type UserService struct{}

func (s UserService) AddAccountIfExists(email string, name string, avatar string) error {
	aOps := user.AccountOps{}
	if isExist, err := aOps.CheckExists(email); err != nil {
		return err
	} else {
		if isExist {
			return nil
		} else {
			return aOps.Init(avatar, name, email)
		}
	}
}

func (s UserService) GetAccountByEmail(email string) (user2.Account, error) {

}

func (s UserService) UpdateUserAccountName(name string, email string) error {

}

func (s UserService) UpdateUserAccountAvatarURL(avatar string, email string) error {

}

func (s UserService) DeleteUserAccount(email string) error {

}

func (s UserService) UpdateUserDetailPhone(phone string, email string) error {

}

func (s UserService) UpdateUserDetailAddress(address string, email string) error {

}
