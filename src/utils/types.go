package utils

import (
	"strings"
	"webserver/src/errors"
)

const IP_BYTES int = 4

type IPAddr struct {
	raw 	[IP_BYTES]byte
	str 	string
	num 	int32
}

func (ip *IPAddr) initIPAddr () {
	for i := 0; i < IP_BYTES; i++ {
		ip.raw[i] = 0
	}
	ip.str = ""
	ip.num = 0
}

func (ip *IPAddr) ParseIP (ip_str string) errors.Error {
	segments := strings.Split(ip_str, ".") 
	if len(segments) != 4 {
		return errors.New("Only IPv4 is supported!")
	}

	return nil
}

type Address struct {
	IP string
}