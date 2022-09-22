package errors

import (
	"errors"
	"fmt"
)

const (
	ErrDB             = errors.New("database failed to process")
	ErrRecordNotFound = Err("record not found")
)

type Err string

func (e Err) Error() string {
	return string(e)
}

func NewError(err error, cause string) error {
	return fmt.Errorf("%w: %s", err, cause)
}

func NewErrorf(err error, args ...interface{}) error {
	return fmt.Errorf(err.Error(), args...)
}
