package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Json( ctx * fasthttp.RequestCtx ){
        ctx.SetContentType( "application/json; charset=utf-8" )
        fmt.Fprintf( ctx, `{"ip":"%s"}`, GetIP( ctx ) )
}
