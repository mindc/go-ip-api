package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// JSON creates response with JSON formatted string
func JSON(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf-8")
	fmt.Fprintf(ctx, `{"ip":"%s"}`, GetIP(ctx))
}
