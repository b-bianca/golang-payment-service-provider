package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repository struct {
	Client       *Client
	Transactions *Transactions
	Payables     *Payables
	Conn         *gorm.DB
	connSQL      *sql.DB
}

func NewRepository(dialector gorm.Dialector) (r *Repository, err error) {
	gdb, err := gorm.Open(dialector)
	if err != nil {
	}

	conn, err := gdb.DB()
	if err != nil {
	}

	return &Repository{
		Client: NewClient(gdb),
		//Transactions: ,
		//Payables: ,
		Conn:    gdb,
		connSQL: conn,
	}, err
}
