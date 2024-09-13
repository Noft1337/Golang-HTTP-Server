package utils

import (
	"errors"
	"fmt"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	"xhttp/src/logging"
)


var log *logging.Logger
const OS string = runtime.GOOS

func init() {
	log = logging.NewLogger(0, logging.LOG_NAME)
}

type UserType struct {
	User 		*user.User
	IsAdmin	bool
}

func isAdmin (user *user.User) (bool, error) {
   gids, err := user.GroupIds()
   if err != nil {
	   return false, err
   }

   for _, g := range gids {
	   if g == "Admin" {
		   return true, nil
	   }
   }

   return false, nil 
}

func getUser() (UserType, error) {
   u := UserType{}
   
   userCur, err := user.Current()
   if err != nil {
	   return u, fmt.Errorf("Port: os.user.Current returned error: %w", err)
   }

   u.User = userCur
   admin, err := isAdmin(u.User)
   if err != nil {
	   u.IsAdmin = false
	   return u, nil
   }

   u.IsAdmin = admin
   return u, nil
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

var RESERVED_PORTS = []int{20,21,22,23,25,53,67,68,110,119,123,135,137,138,139,143,161,162,179,194,445,465,587,993,995,1433,3306,3389,5432}
const SYSTEM_PRIVATE_BEGIN_PORT int = 49152
const SYSTEM_PORT_MAX			int = 65535

func ValidatePort(port int) error {
	if port < 1 || port >  SYSTEM_PORT_MAX {
		return errors.New("Port: must be in range 1 - 65535")
	}

	u, err := getUser()
	if err != nil {
		return fmt.Errorf("Port: utils.getUser returned error: %w", err)
	}

	if OS == "windows" || OS == "linux" {
		for _, r := range RESERVED_PORTS {
			if port == r {
				return fmt.Errorf("Port: %d is a reserved system port", port)
			}
		}

		if port >= SYSTEM_PRIVATE_BEGIN_PORT{
			log.Warn("Port: All ports above %d are reserved for the system to use", SYSTEM_PRIVATE_BEGIN_PORT)
		}

		if port <= 1024 && !u.IsAdmin {
			return fmt.Errorf("Port: Ports below 1024 need to be opened with Administrator Privilege. Permission denied")
		}
	} else {
		return errors.New("Module suupports only Windows/Linux at the time.")
	}
	
	return nil
}
