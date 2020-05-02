# Simple RemoteIP REST API

Simple REST API to get your remote ip address.  
Standalone http-server written in [GO](https://golang.org/)

## Status

[https://ip.mindc.net](https://ip.mindc.net) running...  
[https://ip4.mindc.net](https://ip4.mindc.net) running...  
[https://ip6.mindc.net](https://ip6.mindc.net) running...  

## API

### [https://ip.mindc.net/plain](https://ip.mindc.net/plain)

method: GET,POST  
response (text/plain):

    8.8.8.8


### [https://ip.mindc.net/json](https://ip.mindc.net/json)

method: GET,POST  
response (application/json):

    {"ip":"8.8.8.8"}

### [https://ip.mindc.net/jsonp?callback=YOUR_CALLBACK](https://ip.mindc.net/jsonp?callback=YOUR_CALLBACK)

method: GET  
response (text/javascript):

    YOUR_CALLBACK("8.8.8.8");

### https://ip.mindc.net/jsonrpc

method: POST  
response (application/json):

    {"jsonrpc":"2.0","id":"xFrB","result":"8.8.8.8"}

_* require valid [JSON-RPC 2.0](http://www.jsonrpc.org/specification) request with any string value as `method`_

## Source code

[https://github.com/mindc/remoteip](https://github.com/mindc/remoteip)

Using [https://github.com/valyala/fasthttp](https://github.com/valyala/fasthttp) as net/http replacement.  
Using [https://github.com/buaazp/fasthttprouter](https://github.com/buaazp/fasthttprouter) as HTTP request router.  
Using [https://github.com/caarlos0/env](https://github.com/caarlos0/env) as ENV parser.  
Using [https://github.com/sindresorhus/github-markdown-css](https://github.com/sindresorhus/github-markdown-css) for Markdown CSS  

2017-2020
