package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// simple text response
func Plain(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, GetIP(ctx))
}
