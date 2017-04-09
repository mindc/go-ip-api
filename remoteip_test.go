package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var apiTests = []struct {
	method   string // GET,POST
	postData string // post data
	url      string
	status   int
	expected string
}{
	//plain
	{`GET`, ``, `http://127.0.0.1:8080/plain`, 200, `127.0.0.1`},
	{`POST`, ``, `http://127.0.0.1:8080/plain`, 200, `127.0.0.1`},
	//json
	{`GET`, ``, `http://127.0.0.1:8080/json`, 200, `{"ip":"127.0.0.1"}`},
	{`POST`, ``, `http://127.0.0.1:8080/json`, 200, `{"ip":"127.0.0.1"}`},
	//josnp
	{`POST`, ``, `http://127.0.0.1:8080/jsonp`, 405, `Method Not Allowed`},

	{`GET`, ``, `http://127.0.0.1:8080/jsonp`, 200, `callback("127.0.0.1");`},
	{`GET`, ``, `http://127.0.0.1:8080/jsonp?call`, 200, `callback("127.0.0.1");`},
	{`GET`, ``, `http://127.0.0.1:8080/jsonp?callback=`, 200, `callback("127.0.0.1");`},
	{`GET`, ``, `http://127.0.0.1:8080/jsonp?callback=G`, 200, `G("127.0.0.1");`},

	//jsonrpc
	//ERROR
	// wrong method
	{`GET`, ``, `http://127.0.0.1:8080/jsonrpc`, 405, `Method Not Allowed`},

	// not valid json
	{`POST`, ``, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error"}}`},
	{`POST`, `{"json`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error"}}`},

	// `id` as empty string
	{`POST`, `{"jsonrpc":"2.0","id":"","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"Invalid Request"}}`},

	// `method` missing
	{`POST`, `{"jsonrpc":"2.0","id":"xDcF"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
	// `method` mixed case
	{`POST`, `{"jsonrpc":"2.0","id":"xDcF","Method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
	// `method` as empty string
	{`POST`, `{"jsonrpc":"2.0","id":"xDcF","method":""}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},

	// `jsonrpc` missing
	{`POST`, `{"id":"xDcF","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
	// `jsonrpc` mixed case
	{`POST`, `{"jsonRPC":"2.0","id":"xDcF","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
	// `jsonrpc` wrong version
	{`POST`, `{"jsonrpc":"2.5","id":"xDcF","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
	// `jsonrpc` not string
	{`POST`, `{"jsonrpc":2.0,"id":"xDcF","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},

	//OK
	// `id` as string
	{`POST`, `{"jsonrpc":"2.0","id":"xDcF","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","result":"127.0.0.1"}`},
	// `id` as int
	{`POST`, `{"jsonrpc":"2.0","id":453,"method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":453,"result":"127.0.0.1"}`},
	// `id` as float
	{`POST`, `{"jsonrpc":"2.0","id":453.4,"method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":453.4,"result":"127.0.0.1"}`},
	// `id` is null
	{`POST`, `{"jsonrpc":"2.0","id":null,"method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, `{"jsonrpc":"2.0","id":null,"result":"127.0.0.1"}`},
	// notification
	{`POST`, `{"jsonrpc":"2.0","method":"GET"}`, `http://127.0.0.1:8080/jsonrpc`, 200, ``},
}

func TestAPI(t *testing.T) {
	hc := http.Client{}
	for _, tt := range apiTests {
		t.Log(tt.method, tt.postData, tt.url)
		req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.postData))
		if err != nil {
			t.Error(err)
		}

		resp, err := hc.Do(req)

		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != tt.status {
			t.Error("expected status:", tt.status, "got:", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		if tt.expected != string(body) {

			t.Error("expected:", tt.expected, "got:", string(body))
		}

	}
}
