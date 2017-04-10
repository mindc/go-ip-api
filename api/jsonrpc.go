package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
func JSONRPCDecode(v interface{}) *JSONRPCResponse {
	var r JSONRPCResponse

	jsonrpcMissing := true
	methodMissing := true

	switch vv := v.(type) {
	case map[string]interface{}:
		for key, value := range vv {
			switch key {
			case "id":
				switch vv := value.(type) {
				case string:
					if len(vv) > 0 {
						id := `"` + vv + `"`
						r.ID = &id
					} else {
						id := `null`
						r.ID = &id
						r.Code = -32600
					}
				case float64:
					id := strconv.FormatFloat(vv, 'f', -1, 64)
					r.ID = &id
				case nil:
					id := `null`
					r.ID = &id
				default:
					id := `null`
					r.ID = &id
					r.Code = -32600

				}
			case "jsonrpc":
				jsonrpcMissing = false
				switch vv := value.(type) {
				case string:
					if vv != "2.0" {
						r.Code = -32600
					}
				default:
					r.Code = -32600
				}

			case "method":
				methodMissing = false
				switch vv := value.(type) {
				case string:
					if len(vv) == 0 {
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
		id := `null`
		r.ID = &id
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

	var v interface{}

	if err := json.Unmarshal(ctx.PostBody(), &v); err != nil {
		fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error"}}`)
		return
	}

	switch vv := v.(type) {
	case map[string]interface{}:
		resp := JSONRPCDecode(vv)
		if resp.ID != nil {
			if resp.Code == 0 {
				fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%s,"result":"%s"}`, *resp.ID, GetIP(ctx))
			} else {
				fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":%s,"error":{"code":%d,"message":"%s"}}`, *resp.ID, resp.Code, JSONRPCErrorCodes[resp.Code])
			}
		}
	case []interface{}:
		if len(vv) == 0 {
			fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"Invalid Request"}}`)
			return
		}

		var r []string
		for _, value := range vv {
			resp := JSONRPCDecode(value)
			if resp.ID != nil {
				if resp.Code == 0 {
					r = append(r, fmt.Sprintf(`{"jsonrpc":"2.0","id":%s,"result":"%s"}`, *resp.ID, GetIP(ctx)))
				} else {
					r = append(r, fmt.Sprintf(`{"jsonrpc":"2.0","id":%s,"error":{"code":%d,"message":"%s"}}`, *resp.ID, resp.Code, JSONRPCErrorCodes[resp.Code]))
				}
			}
		}

		if len(r) > 0 {
			ctx.WriteString("[" + strings.Join(r, ",") + "]")
		}

	default:
		fmt.Fprintf(ctx, `{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"Invalid Request"}}`)
	}
}
