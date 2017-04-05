package main

import (
        "fmt"
        "github.com/julienschmidt/httprouter"
	"github.com/mindc/go-ip-api/api"
	"github.com/caarlos0/env"
        "log"
        "net/http"
	"io/ioutil"
)

var (
    Version string
    BUILD string
)

var htdoc string

func Start( w http.ResponseWriter, r *http.Request, _ httprouter.Params ){
	fmt.Fprint( w, htdoc )
}

type Config struct {
    Port	string	`env:"PORT" envDefault:"8080"`
    PortSSL	string	`env:"PORT_SSL"`
    Gzip	bool	`env:"GZIP" envDefault:"false"`
    HtDoc	string	`env:"HTDOC"`
    SSLCert	string	`env:"SSL_CERT"`
    SSLKey	string	`env:"SSL_KEY"`
}


func main() {
    fmt.Printf( "Version: %s\n", Version )
    fmt.Printf( "Build: %s\n", BUILD )

    //read confgiuration from ENV
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
	    fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", cfg)

    //read default htdoc
	if cfg.HtDoc != "" {
            fh, err := ioutil.ReadFile( cfg.HtDoc )
	    if err != nil {
		log.Fatal( err )
	    }
	    htdoc = string( fh )
	} else {
	    htdoc = "It works!"
	}

    //configure router
        router := httprouter.New()

        router.GET( "/", Start )

        router.GET( "/plain", api.Plain )
        router.POST( "/plain", api.Plain )

        router.GET( "/json", api.Json )
        router.POST( "/json", api.Json )

        router.GET( "/jsonp", api.Jsonp )

	router.POST( "/jsonrpc", api.Jsonrpc )

    //main loops
	if cfg.PortSSL != "" {
    	    go func() {
		log.Fatal( http.ListenAndServeTLS( ":" + cfg.PortSSL, cfg.SSLCert, cfg.SSLKey, router ) )
	    }()
	}

        log.Fatal( http.ListenAndServe( ":" + cfg.Port, router ) )


}

