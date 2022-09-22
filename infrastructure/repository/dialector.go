package repository

import (
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres(u *url.URL) gorm.Dialector {
	return postgres.Open(u.String())
}
