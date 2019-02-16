package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cholaraja/revise/goweb/basicStructure/internal/products"
	"github.com/jmoiron/sqlx"
)

// Product defines all the handlers related to products.
// It holds application state needed by the handler methods
type Product struct {
	DB *sqlx.DB
}

// ListProducts uses List from products to encode data
func (p *Product) ListProducts(w http.ResponseWriter, r *http.Request) {
	list, err := products.List(p.DB)
	if err != nil {
		log.Fatalf("error: selecting products: %s", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(list)
}
