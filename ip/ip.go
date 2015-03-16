package ip

import (
	"net"
	"net/http"
	"strings"
)

// Client will attempt to return one and only one
// ip address given a full HTTP Request Object.
// It will also strip ports from IPs if one exists.
func Client(req *http.Request) string {
	addr := req.Header.Get("X-Real-IP")
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = req.RemoteAddr
		}
	}

	// Due to load balancers and CDN's IPs are often represented as
	// ip1, ip2, ip3
	first := strings.Split(addr, ",")[0]

	// Sometimes an IP with a port can get through. So we trim it.
	ip, _, err := net.SplitHostPort(first)
	if err != nil {
		return cleanIPv6(first)
	}

	return ip
}

func cleanIPv6(addr string) string {
	addr = strings.Trim(addr, "[")
	addr = strings.Trim(addr, "]")

	return addr
}
