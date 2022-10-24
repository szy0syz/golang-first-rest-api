package db

import (
	"database/sql"
	"errors"
	"fmt"
)

func HandleError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrObjectNotFound{}
	}
	return err
}

// ErrObjectNotFound is used to indicate that selecting an individual object
// yielded no result. Declared as type, not value, for consistency reasons.
type ErrObjectNotFound struct{}

//TODO
/*
* 如果不写下面的实现就报这样的错误，暂时没搞明白
* cannot use (ErrObjectNotFound literal) (value of type ErrObjectNotFound)
* as error value in return statement:
* ErrObjectNotFound does not implement error (missing method Error)
 */

func (ErrObjectNotFound) Error() string {
	return "object not found"
}

func (ErrObjectNotFound) Unwrap() error {
	return fmt.Errorf("object not found")
}
