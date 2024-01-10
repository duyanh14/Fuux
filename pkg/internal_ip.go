package pkg

import (
	"errors"
	"net"
)

var internalAddrs *[]net.Addr

func InternalIP(ipAddress string) (bool, error) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return false, errors.New("invalid ip address")
	}

	if internalAddrs == nil {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return false, nil
		}
		internalAddrs = &addrs
	}

	for _, addr := range *internalAddrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.Contains(ip) {
				return true, nil
			}
		}
	}

	return false, nil
}
