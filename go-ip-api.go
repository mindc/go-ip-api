package main

import (
        "fmt"
        "github.com/julienschmidt/httprouter"
        "log"
        "net/http"
        "net"
	"encoding/json"
	"io/ioutil"
)

var index []byte

func GetIP ( r * http.Request ) ( string ) {
        ipaddr, _, _ := net.SplitHostPort( r.RemoteAddr )
	return ipaddr
}

func Start( w http.ResponseWriter, r *http.Request, _ httprouter.Params ){
	fmt.Fprint( w, string( index ) )
}

func Plain( w http.ResponseWriter, r *http.Request, _ httprouter.Params ){
        fmt.Fprint( w, GetIP( r ) )
}

func Json( w http.ResponseWriter, r * http.Request, _ httprouter.Params ){
        w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
        fmt.Fprintf( w, "{\"ip\":\"%s\"}", GetIP( r ) )
}

func Jsonp( w http.ResponseWriter, r * http.Request, p httprouter.Params ){
        w.Header().Set( "Content-Type", "text/javascript; charset=utf-8" )

	err := r.ParseForm()
	if err != nil {
	}

	callback := "callback"

	if val, ok := r.Form["callback"]; ok && len( val[ len(val)-1 ] ) > 0 {
	    callback = val[ len(val)-1 ]
	}

        fmt.Fprintf( w, "%s(\"%s\");", callback, GetIP( r ) )
}

type JSONRPC struct {
    ID * string `json:"id,omitempty"`
    Jsonrpc * string `json:"jsonrpc,omitempty"`
    Method * string `json:"method,omitempty"`
}


func Jsonrpc( w http.ResponseWriter, r * http.Request, p httprouter.Params ){
        w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
	w.Header().Set( "Access-Control-Allow-Origin", r.Header.Get( "Origin" ) )
	w.Header().Set( "Access-Control-Allow-Credentials", "true" )



	var j JSONRPC

	if r.Body == nil {
	    http.Error( w, "PSRB", 400 )
	    return
	}

	err := json.NewDecoder( r.Body ).Decode( &j )

	if err != nil {
	    fmt.Fprintf( w, `{"jsonrpc":"2.0","error":{"message":"%s"}}`, err.Error() )
	    return
	}

	if j.ID == nil || j.Jsonrpc == nil || *j.Jsonrpc != "2.0" || j.Method == nil {
	    if j.ID == nil {
    		fmt.Fprintf( w, `{"jsonrpc":"2.0","error":{"message":"Invalid Request"}}` )
	    } else {
    		fmt.Fprintf( w, `{"jsonrpc":"2.0","id":"%s","error":{"message":"Invalid Request"}}`, *j.ID )
	    }
	    return
	}


        fmt.Fprintf( w, `{"jsonrpc":"2.0","id":"%s","result":"%s"}`, *j.ID, GetIP( r ) )
}




func main() {
	var err error
        index, err = ioutil.ReadFile( "/root/gocode/index.html" )
	if err != nil {
	    log.Panic( err.Error() )
        }

        router := httprouter.New()

        router.GET( "/", Start )

        router.GET( "/plain", Plain )
        router.POST( "/plain", Plain )

        router.GET( "/json", Json )
        router.POST( "/json", Json )

        router.GET( "/jsonp", Jsonp )

	router.POST( "/jsonrpc", Jsonrpc )

        log.Fatal( http.ListenAndServe( ":80", router ) )
}

