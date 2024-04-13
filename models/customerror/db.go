package customerror

import "errors"

type DBError struct {
}

func (e *DBError) ConnectError(mess string) error {
	return errors.New("error when connect to database | message: { " + mess + " } ")
}
func (e *DBError) QueryError(nameQuery string, mess string) error {
	return errors.New("error when execute " + nameQuery + " | message: { " + mess + " } ")
}
func (e *DBError) RecordNotFound(name string) error {
	return errors.New("error record not found name: { " + name + " } ")
}
