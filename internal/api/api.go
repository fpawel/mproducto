package api

import (
	"github.com/fpawel/mproducto/internal/api/restapi"
	"github.com/fpawel/mproducto/internal/api/restapi/op"
	"github.com/fpawel/mproducto/internal/app"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/powerman/structlog"
	"net/http"
)

type service struct {
	log *structlog.Logger
	app app.App
}

// Config contains configuration for internal API service.
type Config struct {
	Host string
	Port int
}

// Serve listens on the TCP network address cfg.Host:cfg.Port and
// handle requests on incoming connections.
func Serve(log *structlog.Logger, application app.App, cfg Config) error {
	svc := &service{
		log: log,
		app: application,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := op.NewMproductoAPI(swaggerSpec)
	api.Logger = log.Printf
	api.APIKeyAuth = svc.authenticate
	api.APIAuthorizer = runtime.AuthorizerFunc(svc.authorize)
	api.GetUserHandler = op.GetUserHandlerFunc(svc.getUser)
	api.PostLoginHandler = op.PostLoginHandlerFunc(svc.postLogin)
	api.PutUserHandler = op.PutUserHandlerFunc(svc.putUser)
	api.GetCatalogueHandler = op.GetCatalogueHandlerFunc(svc.getCatalogue)
	api.GetProductsHandler = op.GetProductsHandlerFunc(svc.getProducts)

	server := restapi.NewServer(api)
	defer log.WarnIfFail(server.Shutdown)

	server.Host = cfg.Host
	server.Port = cfg.Port

	// The middleware executes before anything.
	globalMiddlewares := func(handler http.Handler) http.Handler {

		logger := makeLogger(swaggerSpec.BasePath())
		fileServer := fileServer(swaggerSpec.BasePath())
		return fileServer(logger(recovery(handleCORS(handler))))
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		accesslog := makeAccessLog(swaggerSpec.BasePath())
		reauthenticate := svc.reauthenticate()
		return reauthenticate( accesslog(handler) )
		//return accesslog(handler)
	}
	server.SetHandler(globalMiddlewares(api.Serve(middlewares)))

	log.Info("protocol", "version", swaggerSpec.Spec().Info.Version,
		"path", swaggerSpec.BasePath())
	return server.Serve()
}
