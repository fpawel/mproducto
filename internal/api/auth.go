package api

import (
	"github.com/fpawel/mproducto/internal/data"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	DB *sqlx.DB
}

func (svc *Auth) Login(cred data.Cred, token *string) error {
	username, err := data.VerifyCredentials(svc.DB, cred)
	if err != nil {
		return err
	}
	*token, err = jwtTokenizeUsername(username)
	return err
}

func (svc *Auth) Register(user data.User, token *string) error {
	err := data.RegisterUser(svc.DB, user)
	if err != nil {
		return err
	}
	*token, err = jwtTokenizeUsername(user.Name)
	return err
}
