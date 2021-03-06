// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/fpawel/mproducto/internal/api/restapi/op"
	"github.com/fpawel/mproducto/internal/app"
)

//go:generate swagger generate server --target ..\..\api --name Mproducto --spec ..\..\..\swagger.yml --api-package op --model-package model --principal app.Auth --exclude-main

func configureFlags(api *op.MproductoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *op.MproductoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "API-Key" header is set
	api.APIKeyAuth = func(token string) (*app.Auth, error) {
		return nil, errors.NotImplemented("api key auth (api_key) API-Key from header param [API-Key] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.GetCatalogueHandler == nil {
		api.GetCatalogueHandler = op.GetCatalogueHandlerFunc(func(params op.GetCatalogueParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetCatalogue has not yet been implemented")
		})
	}
	if api.GetProductsHandler == nil {
		api.GetProductsHandler = op.GetProductsHandlerFunc(func(params op.GetProductsParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetProducts has not yet been implemented")
		})
	}
	if api.GetUserHandler == nil {
		api.GetUserHandler = op.GetUserHandlerFunc(func(params op.GetUserParams, principal *app.Auth) middleware.Responder {
			return middleware.NotImplemented("operation .GetUser has not yet been implemented")
		})
	}
	if api.PostLoginHandler == nil {
		api.PostLoginHandler = op.PostLoginHandlerFunc(func(params op.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation .PostLogin has not yet been implemented")
		})
	}
	if api.PutUserHandler == nil {
		api.PutUserHandler = op.PutUserHandlerFunc(func(params op.PutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation .PutUser has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
