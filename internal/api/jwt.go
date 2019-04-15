package api

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
type usernameClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func jwtParseUsername(tokenString string) (string, error) {
	claims := new(usernameClaims)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", ErrUnauthorized
	}
	return claims.Username, nil
}

func jwtTokenizeUsername(userName string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &usernameClaims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
