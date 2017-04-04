#!/bin/bash


PORT=${PORT:-80}

echo default page
curl -f http://ip.mindc.net:$PORT
echo

echo 404
curl -f http://ip.mindc.net:$PORT/404
echo

echo plain GET
curl -f http://ip.mindc.net:$PORT/plain
echo
echo plain POST
curl -f http://ip.mindc.net:$PORT/plain
echo

echo json GET
curl -f http://ip.mindc.net:$PORT/json
echo
echo json POST
curl -f --data "" http://ip.mindc.net:$PORT/json
echo

echo jsonp GET
curl -f http://ip.mindc.net:$PORT/jsonp
echo
curl -f http://ip.mindc.net:$PORT/jsonp?call
echo
curl -f http://ip.mindc.net:$PORT/jsonp?callback=
echo
curl -f http://ip.mindc.net:$PORT/jsonp?callback=I
echo

echo jsonp POST
curl -f --data "" http://ip.mindc.net:$PORT/jsonp
echo


echo jsonrpc POST
curl -f --data '{"jsonrpc":"2.0","id":"1","method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo
curl -f --data '{"jsonrpc":"2.0","id":1,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo
curl -f --data '{"jsonrpc":"2.0","method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo


echo



