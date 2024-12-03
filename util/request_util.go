package util

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
)

func GetIpFromRequest(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		// If X-Forwarded-For is empty, fall back to r.RemoteAddr
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func GetExpectedParam(expectedKey string, query url.Values) (string, error) {
	key, keyPresent := query[expectedKey]
	if !keyPresent {
		return "", errors.New(fmt.Sprintf("Could not find expected parameter key in query"))
	}
	if key[0] == "" {
		return "", errors.New("Empty param is not a valid argument")
	}
	return key[0], nil
}
