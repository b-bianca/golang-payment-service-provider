package config

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/golang-payment-service-provider/infrastructure/repository"
)

const (
	postgresURL     = "POSTGRES_URL"
	maxOpenConns    = "MAX_OPEN_CONNS"
	maxIdleConns    = "MAX_IDLE_CONNS"
	connMaxLifetime = "CONN_MAX_LIFETIME_SEC"
	connMaxIdleTime = "CONN_MAX_IDLE_SEC"
)

func NewPostgreDB() (repo *repository.Repository) {
	url := GetPostgreURL()

	d := repository.Postgres(url)
	repo, err := repository.NewRepository(d)
	if err != nil {
		log.Panicf("failed to create postgres db: %v", err)
	}

	if v, ok := getMaxOpenConns(); ok {
		repo.SetMaxOpenConns(v)
	}

	if v, ok := getMaxIdleConns(); ok {
		repo.SetMaxIdleConns(v)
	}

	if v, ok := getConnMaxLifetime(); ok {
		repo.SetConnMaxLifetime(v)
	}

	if v, ok := getConnMaxIdleTime(); ok {
		repo.SetConnMaxIdleTime(v)
	}
	return
}

func GetPostgreURL() *url.URL {
	u := os.Getenv(postgresURL)
	if u == "" {
		log.Panicf("%s var must not be empty", postgresURL)
	}

	url, err := url.Parse(u)
	if err != nil {
		log.Panicf("%s cannot be parsed", postgresURL)
	}

	return url
}

func getMaxOpenConns() (int, bool) {
	v := os.Getenv(maxOpenConns)

	iValue, err := strconv.Atoi(v)
	if err != nil {
		log.Print("maxOpenConns not found. Using default.")
		return 0, false
	}

	return iValue, true
}

func getMaxIdleConns() (int, bool) {
	v := os.Getenv(maxIdleConns)

	iValue, err := strconv.Atoi(v)
	if err != nil {
		log.Print("maxIdleConns not found. Using default.")
		return 0, false
	}

	return iValue, true
}

func getConnMaxLifetime() (time.Duration, bool) {
	v := os.Getenv(connMaxLifetime)

	if v == "" {
		log.Print("connMaxLifetime not found. Using default.")
		return 0, false
	}

	t, err := time.ParseDuration(v + "s")
	if err != nil {
		log.Panicf("reading action not properly configured")
	}
	return t, true
}

func getConnMaxIdleTime() (time.Duration, bool) {
	v := os.Getenv(connMaxIdleTime)

	if v == "" {
		log.Print("connMaxIdleTime not found. Using default.")
		return 0, false
	}

	t, err := time.ParseDuration(v + "s")
	if err != nil {
		log.Panicf("reading action not properly configured")
	}
	return t, true
}
