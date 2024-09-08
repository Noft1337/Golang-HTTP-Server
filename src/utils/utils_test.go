package utils

import (
	"testing"
)


func TestValidateIPv4(t *testing.T) {
	var err error
	badIPs := []string{ "", "notAnIP", "127.0.1", "192.168.345.1", "ff:ff:ff:ff:ff:ff", "1.1.1.0" }
	goodIPs := []string{ "1.1.1.1",  "127.0.0.1", "192.168.234.123"}

	for _, v := range badIPs {
		err = ValidateIPv4(v)
		if err == nil {
			t.Errorf("ValidateIPv4(): Incorrect IP Validated as Correct: %s", v)	
			t.Fail()
		}
	}

	for _, v := range goodIPs {
		err = ValidateIPv4(v)
		if err != nil {
			t.Errorf("ValidateIPv4(): False Positive for: %s (Error: %v)", v, err)
			t.Fail()
		}
	}
}