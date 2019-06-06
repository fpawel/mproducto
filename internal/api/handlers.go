package api

import (
	"github.com/fpawel/mproducto/internal/api/model"
	"github.com/fpawel/mproducto/internal/api/restapi/op"
	"github.com/fpawel/mproducto/internal/app"
	"github.com/go-openapi/runtime/middleware"
)

func (svc *service) getUser(params op.GetUserParams, auth *app.Auth) middleware.Responder {
	ctx, log := fromRequest(params.HTTPRequest, auth)
	user, err := svc.app.User(ctx, log, auth.Token)
	if err != nil {
		return defError(err, op.NewGetUserDefault(0)).(middleware.Responder)
	}
	return op.NewGetUserOK().WithPayload(&model.User{
		Name:  &user.Name,
		Email: &user.Email,
	})
}
