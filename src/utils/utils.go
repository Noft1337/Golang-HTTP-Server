package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"xhttp/src/logging"
)

var log *logging.Logger

func init() {
	log = logging.NewLogger(0, logging.LOG_NAME)
}

func ValidateIPv4(ip string) error {
	ip = strings.ReplaceAll(ip, " ", "")
	ipLen := len(ip)
	if ipLen > 15 || ipLen < 7 {
		return errors.New("Bad IPv4 Length")
	}

	ipSegments := strings.Split(ip, ".")
	if len(ipSegments) != 4 {
		return errors.New("IPv4 can only have 4 segments")
	}

	for i := 0; i < 4; i++ {
		seg := ipSegments[i]

		segInt, err := strconv.Atoi(seg)
		if err != nil {
			return err
		}

		if i == 3 &&  segInt == 0 {
			return errors.New("IPv4 Last Segment Can't be 0")
		}

		if 0 > segInt || segInt > 255 {
			errMsg := fmt.Sprintf("IPv4 only supports segments between 0 & 254, Segment %d is %d", i, segInt)
			return errors.New(errMsg)
		}
	}

	return nil
}

func ValidateIPv6(ip string) error {
	return errors.New("IPv6 isn't supported yet")
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
