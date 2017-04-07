package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Jsonp( ctx * fasthttp.RequestCtx ){
        ctx.SetContentType( "text/javascript; charset=utf-8" )

        callback := "callback"
        val := ctx.FormValue( "callback" )

	if len( val ) > 0 {
	    callback = string( val )
	}
	
	


        fmt.Fprintf( ctx, `%s("%s");`, callback, GetIP( ctx ) )
}

