package api

import (
	"github.com/fpawel/mproducto/internal/data"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	DB *sqlx.DB
}

type AuthLoginArg struct {
	Name string
	Pass string
}

func (svc *Auth) Login(c AuthLoginArg, token *string) error {
	userProfile, err := data.VerifyCredentials(svc.DB, c.Name, c.Pass)
	if err != nil {
		return err
	}
	*token, err = jwtTokenizeIdPass(userProfile.UserID, userProfile.Pass)
	return err
}

type AuthProfileResult struct {
	Name  string
	Email string
}

func (svc *Auth) Profile(token [1]string, r *AuthProfileResult) error {

	userID, pass, err := jwtParseIdPass(token[0])
	if err != nil {
		return err
	}
	u, err := data.VerifyUserID(svc.DB, userID, pass)
	if err != nil {
		return err
	}
	r.Name = u.Name
	r.Email = u.Email
	return nil
}

type AuthRegisterArg struct {
	Name  string
	Email string
	Pass  string
}

func (svc *Auth) Register(u AuthRegisterArg, token *string) error {
	newUserProfile := &data.UserProfile{
		Name:  u.Name,
		Email: u.Email,
		Pass:  u.Pass,
		Role:  "regular_user",
	}
	err := data.RegisterUser(svc.DB, newUserProfile)
	if err != nil {
		return err
	}
	*token, err = jwtTokenizeIdPass(newUserProfile.UserID, newUserProfile.Pass)
	return err
}

func (svc *Auth) Unregister(token [1]string, _ *struct{}) error {

	userID, pass, err := jwtParseIdPass(token[0])
	if err != nil {
		return err
	}
	_, err = data.VerifyUserID(svc.DB, userID, pass)
	if err != nil {
		return err
	}

	_, err = svc.DB.Exec(`DELETE FROM user_profile WHERE user_id = $1`, userID)
	return err
}
