# go-ip-api

Simple API to get your remote ip address.

## Status

Running...

## (http://ip.mindc.net/plain)

method: GET,POST
response (text/plain):

    8.8.8.8


## (http://ip.mindc.net/json)

method: GET,POST
response (application/json):

    {"ip":"8.8.8.8"}

## (http://ip.mindc.net/jsonp?callback=YOUR_CALLBACK)

method: GET
response (text/javascript):

    YOUR_CALLBACK("8.8.8.8");

## http://ip.mindc.net/jsonrpc

method: POST
response (application/json):

    {"jsonrpc":"2.0","id":"xFrB","result":"8.8.8.8"}

_require valid [JSON-RPC 2.0](http://www.jsonrpc.org/specification) request with any value as `method`_

## Author

(https://github.com/mindc)
