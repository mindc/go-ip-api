#!/bin/bash


PORT=${PORT:-80}

echo  notification/no output
curl -f --data '{"jsonrpc":"2.0","method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo id as string
curl -f --data '{"jsonrpc":"2.0","id":"1","method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo id as int
curl -f --data '{"jsonrpc":"2.0","id":1,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo id is null
curl -f --data '{"jsonrpc":"2.0","id":null,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo wrong jsonrpc
curl -f --data '{"jsonrpc":2.0,"id":null,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo wrong jsonrpc
curl -f --data '{"jsonrpc":"2.1","id":34,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo wrong jsonrpc
curl -f --data '{"jsonpc":"2.0","id":34,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo


echo wrong method
curl -f --data '{"jsonrpc":"2.0","id":34,"1method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo wrong id
curl -f --data '{"jsonrpc":"2.0","idd":34,"method":"getIP"}' http://ip.mindc.net:$PORT/jsonrpc
echo

echo empty
curl -f --data '' http://ip.mindc.net:$PORT/jsonrpc
echo









echo



