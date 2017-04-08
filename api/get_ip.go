package api

import (
	"github.com/valyala/fasthttp"
)

// GetIP return remote ip address as string 
func GetIP(ctx *fasthttp.RequestCtx) string {
	i := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if len(i) > 0 {
	    return i
	}
	return ctx.RemoteIP().String()
}
