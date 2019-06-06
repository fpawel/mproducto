package app

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrWrongCred    = errors.New("не верное имя пользователя или пароль")
	ErrUserExists   = errors.New("пользователь с таким именем уже зарегестрирован")

	// the JWT key used to create the signature
	jwtKey = []byte("my_secret_key")
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type idPassClaims struct {
	ID   int64
	Pass string
	jwt.StandardClaims
}

func jwtParseIdPass(tokenString string) (int64, string, error) {

	claims := new(idPassClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", ErrUnauthorized
	}
	return claims.ID, claims.Pass, nil
}

func jwtTokenizeIdPass(userID int64, pass string) (string, error) {

	claims := &idPassClaims{
		userID,
		pass,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
