package config

import (
	"log"
	"os"

	"github.com/namsral/flag"
)

var Domain string
var Port int
var Logger *log.Logger

// New creates a new app-config instance.
func New() {
	flag.StringVar(&Domain, "domain", "http://localhost:8080", "the server domain")
	flag.IntVar(&Port, "port", 8080, "the server port")
	flag.Bool("logging", true, "enables server logging")
	flag.Parse()

	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}
