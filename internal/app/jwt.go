package app

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
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

func jwtParseUserClaims(tokenString string) (int64, string, error) {

	userClaims := new(userClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", ErrUnauthorized
	}
	return userClaims.ID, userClaims.Pass, nil
}

func jwtTokenizeUserClaims(userID int64, pass string) (string, error) {

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
