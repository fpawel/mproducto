package api

import (
	"github.com/fpawel/mproducto/internal/auth"
)

type Auth struct {
}

type UserInfo struct {
	Username string `json:"username"`
	Auth     bool   `json:"auth"`
}

func (_ *Auth) UserInfo(tokenString [1]string, info *UserInfo) (err error) {
	c, err := auth.ParseUsernameFromJwtToken(tokenString[0])
	if err == auth.ErrUnauthorized {
		return nil
	}
	if err != nil {
		return err
	}
	if c == nil {
		panic("unexpected")
	}
	info.Auth = true
	info.Username = c.Username
	return
}

func (_ *Auth) Login(credentials auth.Credentials, tokenString *string) (err error) {
	*tokenString, err = auth.Login(credentials)
	return
}

func (_ *Auth) New(credentials auth.Credentials, tokenString *string) (err error) {

	*tokenString, err = auth.AddNewUser(credentials)
	return nil
}
