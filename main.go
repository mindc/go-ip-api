package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/buaazp/fasthttprouter"
	"github.com/caarlos0/env"
	"github.com/mindc/go-ip-api/api"
	"github.com/valyala/fasthttp"
)

var (
	Version string // Version
	BUILD   string // BUILD
)

var htdoc string

// Start handling default start page
func Start(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	fmt.Fprint(ctx, htdoc)
}

// Config holds configuration strings
type Config struct {
	port    string `env:"PORT" envDefault:"8080"`
	portSSL string `env:"PORT_SSL"`
	htDoc   string `env:"HTDOC"`
	sslCert string `env:"SSL_CERT"`
	sslKey  string `env:"SSL_KEY"`
}

func main() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Build: %s\n", BUILD)

	//read configuration from ENV
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", cfg)

	//read default htdoc
	if len(cfg.htDoc) > 0 {
		fh, err := ioutil.ReadFile(cfg.htDoc)
		if err != nil {
			log.Fatal(err)
		}
		htdoc = string(fh)
	} else {
		htdoc = "It works!"
	}

	//configure router
	router := fasthttprouter.New()

	router.GET("/", Start)

	router.GET("/plain", api.Plain)
	router.POST("/plain", api.Plain)

	router.GET("/json", api.JSON)
	router.POST("/json", api.JSON)

	router.GET("/jsonp", api.JSONP)

	router.POST("/jsonrpc", api.JSONRPC)

	//main loops
	if len(cfg.portSSL) > 0 {
		tlsConfig := &tls.Config{}
		tlsConfig.Certificates = make([]tls.Certificate, 1)
		tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(cfg.sslCert, cfg.sslKey)
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig.BuildNameToCertificate()
		listener, err := tls.Listen("tcp", ":"+cfg.portSSL, tlsConfig)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			log.Fatal(fasthttp.Serve(listener, router.Handler))
		}()

	}

	listener, err := net.Listen("tcp", ":"+cfg.port)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(fasthttp.Serve(listener, router.Handler))

}
