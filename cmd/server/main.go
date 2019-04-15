package main

import (
	"flag"
	"fmt"
	"github.com/fpawel/mproducto/internal/api"
	"github.com/fpawel/mproducto/internal/data"
	"github.com/gorilla/handlers"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

func main() {

	var (
		host  = flag.String("host", "localhost", "Host to run this service on")
		port  = flag.Int("port", 3001, "Port to run this service on")
		dbCfg data.Config
	)

	flag.IntVar(&dbCfg.Port, "pgport", 5432, "Postgres port")
	flag.StringVar(&dbCfg.Host, "pghost", "localhost", "Postgres host")
	flag.StringVar(&dbCfg.User, "pguser", "postgres", "Postgres user")
	flag.StringVar(&dbCfg.Pass, "pgpass", "", "Postgres password")

	flag.Parse()

	if len(dbCfg.Pass) == 0 {
		log.Fatal("Postgres password must be set")
	}

	db, err := data.NewConnectionDB(dbCfg)
	if err != nil {
		log.Fatal("data base error: ", err)
	}
	defer func() {
		_ = db.Close()
	}()

	// Server export an object of type Auth.
	rpcMustRegister(&api.Auth{db})

	// Server provide a HTTP transport on /rpc endpoint.
	http.Handle("/rpc",
		corsHandler{handlers.LoggingHandler(
			os.Stdout,
			jsonrpc2.HTTPHandler(nil))})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil); err != nil {
		log.Fatal(err)
	}
}

func rpcMustRegister(x interface{}) {
	if err := rpc.Register(x); err != nil {
		panic(err)
	}
}
