package app

import (
	"context"
	"github.com/fpawel/mproducto/internal/api/model"
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
	UserID int64
	Pass   string
}

type User = data.User

// App provides application features service.
type App interface {
	GetUser(Ctx, Log, Auth) (User, error)
	AddNewUser(Ctx, Log, User) (string,error)
	Authenticate(string) (*Auth, error)
	Authorize(Auth) error
	Login(ctx Ctx, log Log, name, pass string) (token string, err error)
	Reauthenticate(apiKey string) (string, error)
	GetProductsByTags(tags []string) (products []*model.Product)
}

// New return new application.
func New(db *sqlx.DB) App {
	return &app{db}
}

type app struct {
	db *sqlx.DB
}

func (app *app) Authenticate(apiKey string) (*Auth, error) {
	userID, pass, err := jwtParseUserClaims(apiKey)
	if err != nil {
		return nil, err
	}
	return &Auth{UserID: userID, Pass: pass}, nil
}

func (app *app) Authorize(auth Auth) error {
	user := data.User{
		UserID: auth.UserID,
		Pass:   auth.Pass,
	}
	return data.GetUserByIDAndPass(app.db, &user)
}

func (app *app) GetUser(_ Ctx, _ Log, auth Auth) (User, error) {
	user := User{
		UserID: auth.UserID,
		Pass:   auth.Pass,
	}
	user.UserID = auth.UserID
	user.Pass = auth.Pass
	if err := data.GetUserByIDAndPass(app.db, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (app *app) AddNewUser(ctx Ctx, log Log, user User) (string, error) {
	user.Role = "regular_user"
	user.UserID = 0
	if err := data.AddNewUser(app.db, &user); err != nil {
		return "", err
	}
	return jwtTokenizeUserClaims(user.UserID, user.Pass)
}

func (app *app) Login(ctx Ctx, log Log, name, pass string) (string, error) {
	user := &data.User{
		Name:  name,
		Email: name,
		Pass:  pass,
	}
	if err := data.GetUserByNameAndPass(app.db, user); err != nil {
		return "", err
	}
	return jwtTokenizeUserClaims(user.UserID, user.Pass)
}

func (app *app) Reauthenticate(apiKey string) (string, error) {

	userID, pass, err := jwtParseUserClaims(apiKey)
	if err != nil {
		return "", err
	}
	return jwtTokenizeUserClaims(userID, pass)
}

func (app *app) GetProductsByTags(tags []string) (products []*model.Product) {

	for _,p := range  data.GetProductsByTags(app.db, tags){
		products = append(products, &model.Product{Name:p.Name, ID:p.ID})
	}
	return
}

var _ App = (*app)(nil)
