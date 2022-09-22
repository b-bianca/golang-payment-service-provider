package entity

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	CPF       string    `gorm:"not null; uniqueIndex"`
	User      string    `gorm:"not null, uniqueIndex"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now();index:option:CONCURRENTLY"`
}

type ClientRepository interface {
	//Create
	//Update
	//Retrieve
	//FetchByIDAndUser
}
