package utils

import (
	"fmt"
	"net"
	"testing"
)


func TestValidateIPv4(t *testing.T) {
	var err error
	badIPs := []string{ "", "notAnIP", "127.0.1", "192.168.345.1", "ff:ff:ff:ff:ff:ff", "1.1.1.0", "121.200.123.42.11" }
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

func TestValidatePort(t *testing.T) {
	var invalidRangePorts = []int {-1, 0, 65537}
	var reservedPorts = 	[]int{21,22,445}
	var privilegedPorts =  	[]int{1, 10, 991}
	var validPorts = 		[]int{3241,8888,50000}

	var err error
	invalidRangeError := "Port: must be in range 1 - 65535"
	reservedPortError := "Port: %d is a reserved system port"
	privilegedPortError := "Port: Ports below 1024 need to be opened with Administrator Privilege. Permission denied"

	for _, p := range invalidRangePorts {
		err = ValidatePort(p)
		if err.Error() != invalidRangeError{
			t.Errorf("Range: %d Failed", p)
			t.Fail()
		}
	}

	for _, p := range reservedPorts {
		err = ValidatePort(p)
		if err.Error() != fmt.Sprintf(reservedPortError, p){
			t.Errorf("Reserved: %d Failed", p)
			t.Fail()
		}
	}

	user, err := getUser()
	if err != nil {
		t.Errorf("Getting User raised: %v", err)
		t.Fail()
	}

	for _, p := range privilegedPorts {
		err = ValidatePort(p)
		if user.IsAdmin {
			if err != nil {
				t.Errorf("Privileged: %d False Positive", p)
			} else {
				if err.Error() != privilegedPortError {
					t.Errorf("Privileged: %d is Validated", p )
					t.Fail()
				}
			}
		}
	}

	for _, p := range validPorts {
		err = ValidatePort(p)
		if err != nil {
			t.Errorf("Valid: False positive with err: %v", err.Error())
			t.Fail()
		}
	}
}

func TestPortAvilable(t *testing.T) {
	somePort := 46612
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", somePort))
	if err != nil {
		t.Logf("Warning, Port %d is used, skipping this test.", somePort)
		t.Skip()
	}

	// Test with a port that is used. 
	b := PortAvailable(somePort)
	if b == true {
		t.Errorf("Port %d shown as available even though it isn't", somePort)
		t.Fail()
	}

	fmt.Printf("3\n")
	err = l.Close() 
	if err != nil {
		t.Logf("Warning, Can't close port %d, skipping test", somePort)
		t.Skip()
	}

	fmt.Printf("4\n")
	// Testing that the port is knowingly available
	b = PortAvailable(somePort)
	if b == false {
		t.Errorf("Port %d shown as unavailable even though it is", somePort)
		t.Fail()
	}
}

func TestRandomizePort(t *testing.T) {
	p := RandomizePort()
	if !PortAvailable(p) {
		t.Errorf("Returned port %d is not available", p)
		t.Fail()
	}

	if p < 1024 || p > SYSTEM_PRIVATE_BEGIN_PORT {
		t.Errorf("Port returned is in invalid range")
		t.Fail()
	}
}