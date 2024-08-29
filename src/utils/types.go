package utils

import (
	"errors"
	"strings"
)

func ValidateIPv4(ip string) error {
	return nil 
}

func ValidateIPv6(ip string) error {
	return nil 
}

// Validates that the IP presented is either IPv4/IPv6 
func ValidateIP(ip string) error {
	if strings.Contains(ip, ".") {
		return ValidateIPv4(ip)
	} else if strings.Contains(ip, ":") {
		return ValidateIPv6(ip)
	}
	return errors.New("Invalid IP format supplied, supporting only IPv4/IPv6")
}