package app

import (
	"context"
	"github.com/fpawel/mproducto/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Log is a synonym for convenience.
type Log = *structlog.Logger

// Auth describes authentication.
type Auth struct {
	Token string
}

type User = data.User

// App provides application features service.
type App interface {
	User(Ctx, Log, string) (User, error)
	AddNewUser(Ctx, Log, *User) error
	Authenticate(string) (*Auth, error)
	Authorize(Auth) error
}

type app struct {
	db *sqlx.DB
}

func (app *app) Authenticate(apiKey string) (*Auth, error) {
	_, err := data.GetUserFromToken(app.db, apiKey)
	if err != nil {
		return nil, err
	}
	return &Auth{Token: apiKey}, nil
}

func (app *app) Authorize(auth Auth) error {
	_, err := data.GetUserFromToken(app.db, auth.Token)
	return err
}

func (app *app) User(ctx Ctx, log Log, token string) (User, error) {
	return data.GetUserFromToken(app.db, token)
}

func (app *app) AddNewUser(ctx Ctx, log Log, user *User) error {
	user.Role = "regular_user"
	return data.RegisterUser(app.db, user)
}

var _ App = (*app)(nil)
