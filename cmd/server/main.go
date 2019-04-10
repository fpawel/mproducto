package main

import (
	"flag"
	"github.com/fpawel/mproducto/internal/api"
	"github.com/gorilla/handlers"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

func main() {

	var port = flag.String("port", "3000", "Port to run this service on")

	flag.Parse()

	// Server export an object of type ExampleSvc.
	rpcMustRegister(&api.Auth{})
	rpcMustRegister(&api.Greet{})

	// Server provide a HTTP transport on /rpc endpoint.
	http.Handle("/rpc",
		corsHandler{handlers.LoggingHandler(
			os.Stdout,
			jsonrpc2.HTTPHandler(nil))})

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func rpcMustRegister(x interface{}) {
	if err := rpc.Register(x); err != nil {
		panic(err)
	}
}
