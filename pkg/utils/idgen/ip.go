package idgen

import (
	"net"
	"os"
)

func ResolveExposedIP() net.IP {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = os.Getenv("HOSTNAME")
	}

	addrList, _ := net.LookupIP(hostname)

	for _, addr := range addrList {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4
		}
	}

	return nil
}
