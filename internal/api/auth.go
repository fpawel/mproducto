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

func (_ *Auth) Register(cred auth.Credentials, result *string) (err error) {

	*result, err = auth.Register(cred)
	return nil
}

func (_ *Auth) ValidateNewUsername(username [1]string, result *string) error {

	err := auth.ValidateNewUsername(username[0])
	if err != nil {
		if err == auth.ErrUserExists {
			*result = err.Error()
			return nil
		}
		return err
	}
	return nil
}
