package repository

import (
	"context"

	"github.com/golang-payment-service-provider/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Client is a layer to abstract operation on Client table
type Client struct {
	db *gorm.DB
}

func NewClient(db *gorm.DB) *Client {
	return &Client{db}
}

func (c *Client) Create(ctx context.Context, cl *entity.Client) (client *entity.Client, err error) {
	dbFn := c.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if create := dbFn.Create(cl); create.Error != nil {
		return nil, handleError(create.Error)
	}
	return cl, nil
}
