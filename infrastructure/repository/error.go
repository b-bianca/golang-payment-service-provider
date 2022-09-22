package repository

import (
	"errors"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"

	errPkg "github.com/b-bianca/golang-payment-service-provider/pkg/errors"
)

var (
	ErrDB             = errors.New("database failed to process")
	ErrRecordNotFound = errors.New("database could not find the record")
)

func handleError(err error) error {
	var pgErr *pgconn.PgError
	switch {
	case errors.As(err, &pgErr):
		return errPkg.NewError(ErrDB, err.Error())
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errPkg.NewError(ErrRecordNotFound, err.Error())
	}
	return err
}
