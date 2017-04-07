package api

import (
	"github.com/valyala/fasthttp"
)

func GetIP ( ctx * fasthttp.RequestCtx ) ( string ) {
	return ctx.RemoteIP().String()
}

