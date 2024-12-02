package util

import (
	"net"
	"net/http"
)

func GetIpFromRequest(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		// If X-Forwarded-For is empty, fall back to r.RemoteAddr
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}
