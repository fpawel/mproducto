package api

import (
	"github.com/fpawel/mproducto/internal/api/model"
	"github.com/go-openapi/swag"
)

type defaultResponder interface {
	SetStatusCode(code int)
	SetPayload(payload *model.Error)
}

func defError(err error, r defaultResponder) defaultResponder {
	r.SetStatusCode(500)
	r.SetPayload(&model.Error{
		Code:    swag.Int32(500),
		Message: swag.String(err.Error()),
	})
	return r
}

func apiError(err error) *model.Error {
	return &model.Error{
		Code:    swag.Int32(500),
		Message: swag.String(err.Error()),
	}
}
