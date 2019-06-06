package api

import (
	"github.com/fpawel/mproducto/internal/app"
	"net/http"
)

func (svc *service) authenticate(apiKey string) (*app.Auth, error) {
	return svc.app.Authenticate(apiKey)
}

func (svc *service) authorize(r *http.Request, principal interface{}) error {
	auth := principal.(*app.Auth)
	return svc.app.Authorize(*auth)
}
