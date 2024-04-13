package tests

import (
	"levi-apis/service"
	"testing"
)

func Test(t *testing.T) {
	uS := service.UserService{}
	if err := uS.AddAccountIfExists("email2@email2", "name", "sdff"); err != nil {
		t.Fatal(err.Error())
	} else {
		t.Log("oke")
	}
}
