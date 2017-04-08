package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// create response in JSONP string
func JSONP(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/javascript; charset=utf-8")

	c := "callback"
	v := ctx.FormValue("callback")

	if len(v) > 0 {
		c = string(v)
	}

	fmt.Fprintf(ctx, `%s("%s");`, c, GetIP(ctx))
}
