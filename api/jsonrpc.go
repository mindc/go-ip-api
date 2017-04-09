package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

// JSONRPCResponse holds basic response
type JSONRPCResponse struct {
	ID   *string
	Code int
}

//  JSONRPCErrorCodes holds code to message map
var JSONRPCErrorCodes = map[int]string{
	-32600: `Invalid Request`,
	-32700: `Parse error`,
}

// JSONRPCDecode decoding JSONRPC request from bytes
func JSONRPCDecode(j []byte) *JSONRPCResponse {
	var r JSONRPCResponse

	var v interface{}
	if err := json.Unmarshal(j, &v); err != nil {
		id := `null`
		r.ID = &id
		r.Code = -32700
		return &r
	}

	jsonrpcMissing := true
	methodMissing := true

	switch vv := v.(type) {
	case map[string]interface{}:
		for key, value := range vv {
			switch key {
			case "id":
				switch vvv := value.(type) {
				case string:
					if len(vvv) > 0 {
						id := `"` + vvv + `"`
						r.ID = &id
					} else {
						id := `null`
						r.ID = &id
						r.Code = -32600
					}
				case float64:
					id := strconv.FormatFloat(vvv, 'f', -1, 64)
					r.ID = &id
				default:
					id := `null`
					r.ID = &id

				}
			case "jsonrpc":
				jsonrpcMissing = false
				switch vvv := value.(type) {
				case string:
					if vvv != "2.0" {
						r.Code = -32600
					}
				default:
					r.Code = -32600
				}

			case "method":
				methodMissing = false
				switch vvv := value.(type) {
				case string:
					if len(vvv) == 0 {
						r.Code = -32600

					}
				default:
					r.Code = -32600
				}
			case "params":

			default:
				r.Code = -32600
			}
		}
	default:
		r.Code = -32600
	}

	if jsonrpcMissing || methodMissing {
		r.Code = -32600
	}

	return &r

}

func JSONRPC(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", string(ctx.Request.Header.Peek("Origin")))
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")

	resp := JSONRPCDecode(ctx.PostBody())

	if resp.ID != nil {
		if resp.Code == 0 {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%s,"result":"%s"}`, *resp.ID, GetIP(ctx))
		} else {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%s,"error":{"code":%d,"message":"%s"}}`, *resp.ID, resp.Code, JSONRPCErrorCodes[resp.Code])
		}
	}

}
