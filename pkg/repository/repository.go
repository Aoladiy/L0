package repository

import "github.com/jmoiron/sqlx"

type Delivery interface {
}

type Item interface {
}

type Order interface {
}

type Payment interface {
}

type Repository struct {
	Delivery
	Item
	Order
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
