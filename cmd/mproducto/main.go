// Example swagger service.
package main

import (
	"flag"
	"fmt"
	"github.com/fpawel/mproducto/internal/api"
	"github.com/fpawel/mproducto/internal/app"
	"github.com/fpawel/mproducto/internal/assets"
	"github.com/fpawel/mproducto/internal/data"
	"github.com/fpawel/mproducto/internal/def"
	"net/http"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//nolint:gochecknoglobals
var (
	// set by ./build
	gitVersion  string
	gitBranch   string
	gitRevision string
	gitDate     string
	buildDate   string

	cmd = strings.TrimSuffix(path.Base(os.Args[0]), ".test")
	ver = strings.Join(strings.Fields(strings.Join([]string{gitVersion, gitBranch, gitRevision, buildDate}, " ")), " ")
	log = structlog.New()
	cfg struct {
		version  bool
		logLevel string
		api      api.Config
		pg       data.PgConfig
	}
)

// Init provides common initialization for both app and tests.
func Init() {
	def.Init()

	flag.BoolVar(&cfg.version, "version", false, "print version")
	flag.StringVar(&cfg.logLevel, "log.level", "debug", "log `level` (debug|info|warn|err)")
	flag.StringVar(&cfg.api.Host, "host", def.Host, "listen on `host`")
	flag.IntVar(&cfg.api.Port, "port", def.Port, "listen on `port` (>0)")

	flag.IntVar(&cfg.pg.Port, "pg-port", 5432, "Postgres port")
	flag.StringVar(&cfg.pg.Host, "pg-host", "localhost", "Postgres host")
	flag.StringVar(&cfg.pg.User, "pg-user", "postgres", "Postgres user")
	flag.StringVar(&cfg.pg.Pass, "pg-pass", "", "Postgres password")

	log.SetDefaultKeyvals(
		structlog.KeyUnit, "main",
	)

	namespace := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(cmd, "_")
	InitMetrics(namespace)
	def.InitMetrics()
	api.InitMetrics(namespace)
}

func main() {



	Init()
	flag.Parse()

	switch {

	case cfg.pg.Pass == "":
		fatalFlagValue("required", "host", cfg.pg.Pass)

	case cfg.api.Host == "":
		fatalFlagValue("required", "host", cfg.api.Host)

	case cfg.api.Port <= 0: // Free nginx doesn't support dynamic ports.
		fatalFlagValue("must be > 0", "port", cfg.api.Port)

	case cfg.version: // Must be checked after all other flags for ease testing.
		fmt.Println(cmd, ver, runtime.Version())
		os.Exit(0)
	}

	// Wrong log.level is not fatal, it will be reported and set to "debug".
	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(cfg.logLevel))
	log.Info("started", "version", ver)

	http.Handle("/metrics", promhttp.Handler())
	go func() { log.Fatal(http.ListenAndServe(cfg.api.Host+":8080", nil)) }()

	db, err := data.NewConnectionDB(cfg.pg)
	if err != nil {
		log.Panic("data base error: ", err)
	}
	defer log.ErrIfFail(db.Close)



	a := app.New(db)

	http.Handle("/", http.FileServer(assets.AssetFS()))
	if err := api.Serve(log, a, cfg.api); err != nil {
		log.Panic(err)
	}
}

// fatalFlagValue report invalid flag values in same way as flag.Parse().
func fatalFlagValue(msg, name string, val interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "invalid value %#v for flag -%s: %s\n", val, name, msg)
	flag.Usage()
	os.Exit(2)
}
