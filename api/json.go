package api

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Json( w http.ResponseWriter, r * http.Request, _ httprouter.Params ){
        w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
        fmt.Fprintf( w, `{"ip":"%s"}`, GetIP( r ) )
}
