package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"sync"
	"time"
)

type Auth struct {
}

type CredentialsContext struct {
	Credentials
	jsonrpc2.Ctx
}

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var (
	users = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
	muUsers sync.Mutex
)

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (_ *Auth) Login(credentials Credentials, tokenString *string) (err error) {
	*tokenString, err = login(credentials)
	return
}

func (_ *Auth) New(credentials Credentials, tokenString *string) (err error) {

	muUsers.Lock()
	if _, exists := users[credentials.Username]; exists {
		muUsers.Unlock()
		err = errors.New("user already exists")
		return
	}
	users[credentials.Username] = credentials.Password
	muUsers.Unlock()

	*tokenString, err = login(credentials)
	return nil
}

func getClaimsFromTokenString(tokenString string) (*Claims, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
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

func login(credentials Credentials) (string, error) {
	// Get the expected password from our in memory map
	muUsers.Lock()
	expectedPassword, ok := users[credentials.Username]
	muUsers.Unlock()

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != credentials.Password {
		return "", ErrUnauthorized
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	// If there is an error in creating the JWT return an internal server error
	return token.SignedString(jwtKey)
}

var ErrUnauthorized = errors.New("unauthorized")
