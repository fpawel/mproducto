package api

import "github.com/fpawel/mproducto/internal/auth"

type Greet struct {
}

func (_ Greet) Hello(tokenString [1]string, reply *string) error {
	claims, err := auth.ParseUsernameFromJwtToken(tokenString[0])
	if err != nil {
		return err
	}
	*reply = "Welcome, " + claims.Username + "!"
	return nil
}
