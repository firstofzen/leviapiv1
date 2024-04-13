package customerror

import "errors"

type ValidateError struct{}

func (v ValidateError) FieldEmpty() error {
	return errors.New("value field is invalid")
}
func (v ValidateError) ValidateStruct(mess string) error {
	return errors.New("Struct hava an error | mess: " + mess)
}
