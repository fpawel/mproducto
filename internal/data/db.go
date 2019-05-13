package data

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type PgConfig struct {
	Port int
	Host,
	User,
	Pass string
}

func NewConnectionDB(c PgConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=mproducto sslmode=disable",
		c.Host, c.Port, c.User, c.Pass)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(conn, "postgres"), nil
}


func DefaultPgConfig() PgConfig{
	return PgConfig{
		Port:5432,
		Host:"localhost",
		User:"postgres",
		Pass:os.Getenv("MPRODUCTO_POSTGRES_PASSWORD"),
	}
}