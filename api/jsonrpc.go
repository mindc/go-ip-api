package api

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type JSONRPC struct {
        ID      json.RawMessage `json:"id,omitempty"`
        Jsonrpc *string         `json:"jsonrpc,omitempty"`
        Method  *string         `json:"method,omitempty"`
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
