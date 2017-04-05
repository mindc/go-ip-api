package api

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Jsonp( w http.ResponseWriter, r * http.Request, p httprouter.Params ){
        w.Header().Set( "Content-Type", "text/javascript; charset=utf-8" )

        err := r.ParseForm()
        if err != nil {
        }

        callback := "callback"

        if val, ok := r.Form["callback"]; ok && len( val[ len(val)-1 ] ) > 0 {
            callback = val[ len(val)-1 ]
        }

        fmt.Fprintf( w, `%s("%s");`, callback, GetIP( r ) )
}

