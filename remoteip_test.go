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
	{`GET`, ``, `http://127.0.0.1/plain`, 200, `127.0.0.1`},
	{`POST`, ``, `http://127.0.0.1/plain`, 200, `127.0.0.1`},

	{`GET`, ``, `http://127.0.0.1/json`, 200, `{"ip":"127.0.0.1"}`},
	{`POST`, ``, `http://127.0.0.1/json`, 200, `{"ip":"127.0.0.1"}`},

	{`POST`, ``, `http://127.0.0.1/jsonp`, 405, `Method Not Allowed`},
	{`GET`, ``, `http://127.0.0.1/jsonp`, 200, `callback("127.0.0.1");`},
	{`GET`, ``, `http://127.0.0.1/jsonp?call`, 200, `callback("127.0.0.1");`},
	{`GET`, ``, `http://127.0.0.1/jsonp?callback=`, 200, `callback("127.0.0.1");`},
	{`GET`, ``, `http://127.0.0.1/jsonp?callback=G`, 200, `G("127.0.0.1");`},

	{`GET`, ``, `http://127.0.0.1/jsonrpc`, 405, `Method Not Allowed`},
	{`POST`, `{"jsonrpc":"2.0","id":"xDcF","method":"GET"}`, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","result":"127.0.0.1"}`},
	{`POST`, `{"jsonrpc":"2.0","id":453,"method":"GET"}`, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":453,"result":"127.0.0.1"}`},
	{`POST`, `{"jsonrpc":"2.0","id":null,"method":"GET"}`, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":null,"result":"127.0.0.1"}`},
	{`POST`, `{"jsonrpc":"2.0","method":"GET"}`, `http://127.0.0.1/jsonrpc`, 200, ``},
	{`POST`, ``, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":null,"error":{"code":-32700,"message":"Parse error","data":"EOF"}}`},
	{`POST`, `{"jsonrpc":2.0,"id":"xDcF"}`, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid request"}}`},

	{`POST`, `{"jsonrpc":2.0,"id":"xDcF","method":"GET"}`, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
	{`POST`, `{"jsonrpc":"2.5","id":"xDcF","method":"GET"}`, `http://127.0.0.1/jsonrpc`, 200, `{"jsonrpc":"2.0","id":"xDcF","error":{"code":-32600,"message":"Invalid Request"}}`},
}

func TestAPI(t *testing.T) {
	hc := http.Client{}
	for _, tt := range apiTests {
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
