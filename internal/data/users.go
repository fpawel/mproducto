package data

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"

)



type UserProfile struct {
	UserID int64  `db:"user_id"`
	Name   string `db:"name"`
	Pass   string `db:"pass"`
	Email  string `db:"email"`
	Role   string `db:"user_role"`
}

var (
	ErrUserExists      = errors.New("пользователь с таким именем уже зарегестрирован")
	ErrWrongCred       = errors.New("не верное имя пользователя или пароль")
	ErrWrongUserIDPass = errors.New("не верный идентификатор пользователя или пароль")
)



func VerifyUserID(db *sqlx.DB, userID int64, pass string) (user UserProfile, err error) {
	err = db.Get(&user,
		`SELECT * FROM user_profile WHERE user_id = $1 AND pass = $2`,
		userID, pass)
	if err == sql.ErrNoRows {
		err = ErrWrongUserIDPass
		return
	}
	return
}

func VerifyCredentials(db *sqlx.DB, name, pass string) (user UserProfile, err error) {
	err = db.Get(&user,
		`SELECT * FROM user_profile WHERE (name = $1 OR email = $1) AND pass = $2`,
		name, pass)
	if err == sql.ErrNoRows {
		err = ErrWrongCred
		return
	}
	return
}

func RegisterUser(db *sqlx.DB, user *UserProfile) error {
	return db.QueryRow(
		`INSERT INTO user_profile(name, email, pass, user_role) VALUES ($1,$2,$3,$4) RETURNING user_id`,
		user.Name, user.Email, user.Pass, user.Role).Scan(&user.UserID)
}
