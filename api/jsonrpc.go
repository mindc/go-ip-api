package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

// JSONRPCRequest holds JSON-RPC 2.0 request
type JSONRPCRequest struct {
	ID      json.RawMessage `json:"id,omitempty"`
	Jsonrpc *string         `json:"jsonrpc,omitempty"`
	Method  *string         `json:"method,omitempty"`
}

// JSONRPC creates valid JSON-RPC 2.0 response
func JSONRPC(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", string(ctx.Request.Header.Peek("Origin")))
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")

	if ctx.PostBody() == nil {
		fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error"}}`)
		return
	}

	var j JSONRPCRequest
	err := json.NewDecoder(strings.NewReader(string(ctx.PostBody()))).Decode(&j)

	if err != nil {
		fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error","data":"%s"}}`, err.Error())
		return
	}

	if j.ID == nil { //notify
		fmt.Fprint(ctx, "")
		return
	}

	var s string
	if err = json.Unmarshal(j.ID, &s); err == nil {
		if len(s) > 0 {
			s = `"` + s + `"`
		} else {
			s = `null`
		}

		if j.Jsonrpc == nil || *j.Jsonrpc != "2.0" || j.Method == nil {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%s,"error":{"code":-32600,"message":"Invalid Request"}}`, s)
		} else {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%s,"result":"%s"}`, s, GetIP(ctx))
		}
		return
	}

	var n uint64
	if err = json.Unmarshal(j.ID, &n); err == nil {
		if j.Jsonrpc == nil || *j.Jsonrpc != "2.0" || j.Method == nil {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%d,"error":{"code":-32600,"message":"Invalid Request"}}`, n)
		} else {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%d,"result":"%s"}`, n, GetIP(ctx))
		}
		return
	}

}
