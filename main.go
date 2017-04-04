package main

import (
        "fmt"
	"flag"
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
    ID	 	json.RawMessage	`json:"id,omitempty"`
    Jsonrpc 	*string 		`json:"jsonrpc,omitempty"`
    Method 	*string 		`json:"method,omitempty"`
}

func Jsonrpc( w http.ResponseWriter, r * http.Request, p httprouter.Params ){
        w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
	w.Header().Set( "Access-Control-Allow-Origin", r.Header.Get( "Origin" ) )
	w.Header().Set( "Access-Control-Allow-Credentials", "true" )

	if r.Body == nil {
	    fmt.Fprintf( w, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error"}}` )
	    return
	}

	var j JSONRPC
	err := json.NewDecoder( r.Body ).Decode( &j )

	if err != nil {
	    fmt.Fprintf( w, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error","data":"%s"}}`, err.Error() )
	    return
	}


	if j.ID == nil { //notify
	    fmt.Fprint( w, "" )
	    return
	}

	var s string
	if err = json.Unmarshal( j.ID, &s ); err == nil {
	    if len ( s ) > 0 {
		s = `"` + s + `"`
	    } else {
		s = `null`
	    }

    	    if j.Jsonrpc == nil || *j.Jsonrpc != "2.0" || j.Method == nil {
		fmt.Fprintf( w, `{"jsonrpc":"2.0","id":%s,"error":{"code":-32600,"message":"Invalid Request"}}`, s )
	    } else {
		fmt.Fprintf( w, `{"jsonrpc":"2.0","id":%s,"result":"%s"}`, s, GetIP( r ) )
	    }
	    return
	}

	var n uint64
	if err = json.Unmarshal( j.ID, &n ); err == nil {
    	    if j.Jsonrpc == nil || *j.Jsonrpc != "2.0" || j.Method == nil {
		fmt.Fprintf( w, `{"jsonrpc":"2.0","id":%d,"error":{"code":-32600,"message":"Invalid Request"}}`, n )
	    } else {
		fmt.Fprintf( w, `{"jsonrpc":"2.0","id":%d,"result":"%s"}`, n, GetIP( r ) )
	    }
	    return
	}

}




func main() {
	var (
		port	= flag.String( "port", "8080", "The server port" )
		htdoc	= flag.String( "htdoc", "index.html", "Default page location" )
	)

	flag.Parse()


	var err error
        index, err = ioutil.ReadFile( *htdoc )
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

        log.Fatal( http.ListenAndServe( ":"+ *port, router ) )
}

