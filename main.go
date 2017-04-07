package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/caarlos0/env"
	"github.com/mindc/go-ip-api/api"
	"github.com/valyala/fasthttp"
)

var (
	Version string
	BUILD   string
)

var htdoc string

func Start(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	fmt.Fprint(ctx, htdoc)
}

type Config struct {
	Port    string `env:"PORT" envDefault:"8080"`
	PortSSL string `env:"PORT_SSL"`
	HtDoc   string `env:"HTDOC"`
	SSLCert string `env:"SSL_CERT"`
	SSLKey  string `env:"SSL_KEY"`
}

func main() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Build: %s\n", BUILD)

	//read confgiuration from ENV
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", cfg)

	//read default htdoc
	if cfg.HtDoc != "" {
		fh, err := ioutil.ReadFile(cfg.HtDoc)
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
	if cfg.PortSSL != "" {
		go func() {
			log.Fatal(fasthttp.ListenAndServeTLS(":"+cfg.PortSSL, cfg.SSLCert, cfg.SSLKey, router.Handler))
		}()
	}

	log.Fatal(fasthttp.ListenAndServe(":"+cfg.Port, router.Handler))

}
