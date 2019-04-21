package main

import (
	"flag"
	"fmt"
	"github.com/fpawel/mproducto/internal/api"
	"github.com/fpawel/mproducto/internal/assets"
	"github.com/fpawel/mproducto/internal/data"
	"github.com/gorilla/handlers"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

func main() {

	var arg struct {
		Host string
		Port int
		Pg   data.PgConfig
	}

	flag.StringVar(&arg.Host, "host", "localhost", "Host to run this service on")
	flag.IntVar(&arg.Port, "port", 3001, "Port to run this service on")
	flag.IntVar(&arg.Pg.Port, "pg-port", 5432, "Postgres port")
	flag.StringVar(&arg.Pg.Host, "pg-host", "localhost", "Postgres host")
	flag.StringVar(&arg.Pg.User, "pg-user", "postgres", "Postgres user")
	flag.StringVar(&arg.Pg.Pass, "pg-pass", "", "Postgres password")

	flag.Parse()

	log.Printf("applied config: %+v\n", arg)

	if len(arg.Pg.Pass) == 0 {
		log.Fatal("Postgres password must be set")
	}

	db, err := data.NewConnectionDB(arg.Pg)
	if err != nil {
		log.Fatal("data base error: ", err)
	}
	defer func() {
		_ = db.Close()
	}()

	// Server export an object of type Auth.
	rpcMustRegister(&api.Auth{db})

	// Server provide a HTTP transport on /rpc endpoint.
	http.Handle("/rpc", corsHandler{
		handlers.LoggingHandler(
			os.Stdout,
			jsonrpc2.HTTPHandler(nil)),
	})

	http.Handle("/", http.StripPrefix(
		"/", http.FileServer(assets.AssetFS())))

	serveURL := fmt.Sprintf("%s:%d", arg.Host, arg.Port)
	log.Printf("serve URL:  http://%s\n", serveURL)
	if err := http.ListenAndServe(serveURL, nil); err != nil {
		log.Fatal(err)
	}
}

func rpcMustRegister(x interface{}) {
	if err := rpc.Register(x); err != nil {
		panic(err)
	}
}
