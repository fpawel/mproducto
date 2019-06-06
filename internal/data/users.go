package data

import (
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
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
	ErrUnauthorized    = errors.New("unauthorized")

	// the JWT key used to create the signature
	jwtKey = []byte("my_secret_key")
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type userClaims struct {
	ID   int64
	Pass string
	jwt.StandardClaims
}

func GetUserByID(db *sqlx.DB, userID int64, pass string) (user User, err error) {
	err = db.Get(&user,
		`SELECT * FROM user_profile WHERE user_id = $1 AND pass = $2`,
		userID, pass)
	if err == sql.ErrNoRows {
		err = ErrWrongUserIDPass
		return
	}
	return
}

func GetUserByNameAndPass(db *sqlx.DB, name, pass string) (user User, err error) {
	err = db.Get(&user,
		`SELECT * FROM user_profile WHERE (name = $1 OR email = $1) AND pass = $2`,
		name, pass)
	if err == sql.ErrNoRows {
		err = ErrWrongCred
		return
	}
	return
}

func RegisterUser(db *sqlx.DB, user *User) error {
	return db.QueryRow(
		`INSERT INTO user_profile(name, email, pass, user_role) VALUES ($1,$2,$3,$4) RETURNING user_id`,
		user.Name, user.Email, user.Pass, user.Role).Scan(&user.UserID)
}

func GetUserFromToken(db *sqlx.DB, tokenString string) (User, error) {
	userClaims := new(userClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return User{}, err
	}
	if !token.Valid {
		return User{}, ErrUnauthorized
	}

	user := User{}

	err = db.Get(&user,
		`SELECT * FROM user_profile WHERE user_id = $1 AND pass = $2`,
		userClaims.ID, userClaims.Pass)
	if err == sql.ErrNoRows {
		return User{}, ErrWrongUserIDPass
	}
	return user, nil
}

func jwtTokenizeIdPass(userID int64, pass string) (string, error) {

	claims := &userClaims{
		userID,
		pass,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
