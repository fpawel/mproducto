package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrUserExists   = errors.New("пользователь с таким именем уже зарегестрирован")

	// the JWT key used to create the signature
	jwtKey = []byte("my_secret_key")

	users = map[string]string{
		"userName1": "password1",
		"userName2": "password2",
	}
)

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type UsernameClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(credentials Credentials) (string, error) {
	if _, exists := users[credentials.Username]; exists {
		return "", ErrUserExists
	}
	users[credentials.Username] = credentials.Password

	return jwtTokenizeUserName(credentials.Username)
}

func Login(credentials Credentials) (string, error) {
	if expectedPassword, ok := users[credentials.Username]; !ok || expectedPassword != credentials.Password {
		return "", ErrUnauthorized
	}
	return jwtTokenizeUserName(credentials.Username)
}

func ValidateNewUsername(username string) error {
	if _, exists := users[username]; exists {
		return ErrUserExists
	}
	return nil
}

func ParseUsernameFromJwtToken(tokenString string) (*UsernameClaims, error) {
	claims := &UsernameClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrUnauthorized
	}
	return claims, nil
}

func jwtTokenizeUserName(userName string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &UsernameClaims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
