package config

import (
	"log"
	"os"

	"github.com/namsral/flag"
)

var Domain string
var Port int
var Logger *log.Logger
var Cors struct {
	Secure bool
	Debug  bool
}

// New creates a new app-config instance.
func New() {
	flag.StringVar(&Domain, "domain", "localhost:8080", "the server domain")
	flag.IntVar(&Port, "port", 8080, "the server port")
	flag.Bool("logging", true, "enables server logging")
	flag.BoolVar(&Cors.Secure, "cors_secure", false, "enables cors in secure mode")
	flag.BoolVar(&Cors.Debug, "cors_debug", false, "enables cors in debug mode")
	flag.Parse()

	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}
