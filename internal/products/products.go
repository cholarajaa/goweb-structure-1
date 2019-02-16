package products

import (
	"github.com/jmoiron/sqlx"
)

// Product struct holds the fields
type Product struct {
	ID       string `db:"product_id"`
	Name     string `db:"name"`
	Price    int    `db:"price"`
	Quantity int    `db:"quantity"`
}

// List function lists all the products
func List(db *sqlx.DB) ([]Product, error) {
	var products []Product
	if err := db.Select(&products, "SELECT * FROM products_new"); err != nil {
		return nil, err
	}
	return products, nil
}
