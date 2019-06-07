package api

import (
	"context"
	"github.com/fpawel/mproducto/internal/api/model"
	"github.com/fpawel/mproducto/internal/api/restapi/op"
	"github.com/fpawel/mproducto/internal/app"
	"github.com/fpawel/mproducto/internal/def"
	"github.com/go-openapi/runtime/middleware"
	"github.com/powerman/structlog"
	"net/http"
)

func (svc *service) getUser(params op.GetUserParams, auth *app.Auth) middleware.Responder {
	ctx, log := svc.fromRequestWithAuth(params.HTTPRequest, auth)
	user, err := svc.app.GetUser(ctx, log, *auth)
	if err != nil {
		return defError(err, op.NewGetUserDefault(0)).(middleware.Responder)
	}
	return op.NewGetUserOK().WithPayload(&model.User{
		Name:  &user.Name,
		Email: &user.Email,
	})
}

func (svc *service) postLogin(params op.PostLoginParams) middleware.Responder {
	ctx, log := svc.fromRequest(params.HTTPRequest)
	token, err := svc.app.Login(ctx, log, *params.Credentials.Name, *params.Credentials.Password)
	if err != nil {
		return defError(err, op.NewPostLoginDefault(0)).(middleware.Responder)
	}
	return op.NewPostLoginCreated().WithPayload(token)
}

func (svc *service) fromRequestWithAuth(r *http.Request, auth *app.Auth) (context.Context, *structlog.Logger) {
	ctx, log := svc.fromRequest(r)
	if auth != nil {
		log = log.SetDefaultKeyvals(def.LogUserID, auth.UserID)
	}
	return ctx, log
}

func (svc *service) fromRequest(r *http.Request) (context.Context, *structlog.Logger) {
	ctx := r.Context()
	log := structlog.FromContext(ctx, nil)
	return ctx, log
}
