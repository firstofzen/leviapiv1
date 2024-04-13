package validate

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"levi-apis/models/customerror"
)

var valErr customerror.ValidateError

type RequestValidate struct{}

func (r RequestValidate) Default(d interface{}) error {
	val := validator.New()
	err := val.Struct(d)
	return errors.New("data have field null or empty | mess: { " + err.Error() + " } ")
}
