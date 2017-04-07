#!/bin/bash

URL=${URL:-http://ip.mindc.net}
URL=${1:-$URL}

echo -e "\e[1mdefault page\e[m"
curl -if ${URL}
echo

echo -e "\e[1mnot found\e[m"
curl -if ${URL}/404
echo

echo -e "\e[1mPlain GET\e[m"
curl -if ${URL}/plain
echo
echo -e "\e[1mPlain POST\e[m"
curl -if ${URL}/plain
echo

echo -e "\e[1mJSON GET\e[m"
curl -if ${URL}/json
echo
echo -e "\e[1mJSON POST\e[m"
curl -if --data "" ${URL}/json
echo

echo -e "\e[1mJSONP GET\e[m"
set -x
curl -if ${URL}/jsonp
curl -if ${URL}/jsonp?call
curl -if ${URL}/jsonp?callback=
curl -if ${URL}/jsonp?callback=I
set +x

echo -e "\e[1mJSONP POST\e[m"
curl -if --data "" ${URL}/jsonp
echo


echo -e "\e[1mJSONRPC POST\e[m"
echo -e "\e[1mID as string\e[m"
curl -if --data '{"jsonrpc":"2.0","id":"1","method":"getIP"}' ${URL}/jsonrpc
echo;echo -e "\e[1mID as number\e[m"
curl -if --data '{"jsonrpc":"2.0","id":1,"method":"getIP"}' ${URL}/jsonrpc
echo;echo -e "\e[1mID is null\e[m"
curl -if --data '{"jsonrpc":"2.0","method":"getIP"}' ${URL}/jsonrpc
echo;echo



