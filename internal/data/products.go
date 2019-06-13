package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/powerman/must"
)

type Product struct {
	ID   int64  `db:"product_id"`
	Name string `db:"name"`
}

func GetProductsByTags(db *sqlx.DB, tags []string) (products []Product) {
	pqTags := pq.Array(tags)
	err := db.Select(&products,
		`SELECT product_id,name FROM product WHERE lower(tags::text)::text[] && lower($1::text)::text[]`,
		pqTags)
	must.AbortIf(err)
	return
}