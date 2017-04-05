package api

import (
	"net/http"
	"net"
)

func GetIP ( r * http.Request ) ( string ) {
        ipaddr, _, _ := net.SplitHostPort( r.RemoteAddr )
        return ipaddr
}

