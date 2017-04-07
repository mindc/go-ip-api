package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Jsonp( ctx * fasthttp.RequestCtx ){
        ctx.SetContentType( "text/javascript; charset=utf-8" )

//        callback := "callback"

        callback := ctx.FormValue( "callback" )

        fmt.Fprintf( ctx, `%s("%s");`, callback, GetIP( ctx ) )
}

