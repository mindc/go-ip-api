package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Plain( ctx * fasthttp.RequestCtx ) {
        fmt.Fprint( ctx, GetIP( ctx ) )
}

