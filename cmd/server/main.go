package main

import (
	"github.com/fpawel/mproducto/internal/api"
	"github.com/gorilla/handlers"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func main() {


	// Server export an object of type ExampleSvc.
	rpcMustRegister(&api.Auth{})

	// Server provide a HTTP transport on /rpc endpoint.
	http.Handle("/rpc",
		corsHandler{handlers.LoggingHandler(
			os.Stdout,
			jsonrpc2.HTTPHandler(nil))})
	lnHTTP, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	defer lnHTTP.Close()
	log.Println("start serving rpc")
	if err := http.Serve(lnHTTP, nil); err != nil {
		log.Panic(err)
	}
}

func rpcMustRegister(x interface{}) {
	if err := rpc.Register(x); err != nil {
		panic(err)
	}
}



type corsHandler struct {
	handler   http.Handler
}

func (h corsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if req.Method == "OPTIONS" {
		return
	}
	h.handler.ServeHTTP(w,req)
}

