package api

import (
	"github.com/valyala/fasthttp"
)

func GetIP(ctx *fasthttp.RequestCtx) string {
	i := string(ctx.Request.Header.Peek("X-Forwarded-For"))

	if i == "" {
		return ctx.RemoteIP().String()
	}
	return i
}
