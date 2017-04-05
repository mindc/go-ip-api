# Simple IP-API

Simple API to get your remote ip address.  
Using [https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) as HTTP request router.  
Using [https://github.com/caarlos0/env]("github.com/caarlos0/env") as ENV parser.

## Status

Running...

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

[https://github.com/mindc/go-ip-api](https://github.com/mindc/go-ip-api)

2017