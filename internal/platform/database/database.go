package database

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver
)

// Open connects to database
func Open() (*sqlx.DB, error) {
	q := url.Values{}
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		Host:     "localhost:5433",
		Path:     "productdb",
		User:     url.UserPassword("postgres", "postgres"),
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}
