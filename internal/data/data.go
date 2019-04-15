package data

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Port int
	Host,
	User,
	Pass string
}

var (
	ErrUserExists = errors.New("пользователь с таким именем уже зарегестрирован")
	ErrWrongCred  = errors.New("не верное имя пользователя или пароль")
)

func NewConnectionDB(c Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=mproducto sslmode=disable",
		c.Host, c.Port, c.User, c.Pass)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(conn, "postgres"), nil
}

func GetCredentialsByName(db *sqlx.DB, name string) (cred Cred, err error) {
	err = db.Get(&cred,
		`SELECT * FROM credential WHERE (name = $1 OR email = $1) AND pass = $2`,
		cred.Name, cred.Pass)
	if err == sql.ErrNoRows {
		err = ErrWrongCred
		return
	}
	return
}

func VerifyCredentials(db *sqlx.DB, cred Cred) (name string, err error) {
	err = db.Get(&name,
		`SELECT name FROM credential WHERE (name = $1 OR email = $1) AND pass = $2`,
		cred.Name, cred.Pass)
	if err == sql.ErrNoRows {
		err = ErrWrongCred
		return
	}
	return
}

func RegisterUser(db *sqlx.DB, user User) error {
	_, err := db.Exec(
		`INSERT INTO credential(name, email, pass, user_role) VALUES ($1,$2,$3,$4)`,
		user.Name, user.Email, user.Pass, user.Role)
	return err
}
