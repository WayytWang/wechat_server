package utils

import (
	"github.com/pkg/errors"
	"net"
)

// ExternalIP 获取公网IP地址
func ExternalIP() (string, error) {
	iFaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iFace.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		adders, err := iFace.Addrs()
		if err != nil {
			err = errors.Wrap(err, "[ExternalIP] [iFace.Addrs()]")
			return "", err
		}
		for _, addr := range adders {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.Wrap(err, "Your machine may not have Internet access")
}

//获取ip
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
