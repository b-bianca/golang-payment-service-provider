package repository_test

import (
	"net/url"
	"testing"

	"github.com/golang-payment-service-provider/infrastructure/repository"
)

func Test_Postgress(t *testing.T) {
	u, err := url.Parse("postgresql://postgres:postgres@localhost:5432/postgres")
	assert.NoError(t, err)

	driver := repository.Postgres(u)
	assert.NotNil(t, driver)
}
