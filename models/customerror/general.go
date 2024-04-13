package customerror

import "errors"

type General struct {
}

func (t General) InvalidFieldName(mess string) error {
	return errors.New("invalid field name for given type | mess: { " + mess + " } ")
}
func (t General) ParseUUID(mess string) error {
	return errors.New("can not parse uuid | mess: { " + mess + " } ")
}
