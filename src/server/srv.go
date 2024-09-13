package server

import (
	"xhttp/src/utils"
)

const (
	CHANNEL_BUFFER_SIZE = 65536
)

type Server struct {
	Address		string
	Port		int
	chCnt		int
	ch 			chan string
}

func NewHTTPServer(addr string, port int) Server {
	
	
	return Server {

	}
}

// This function takes the ip and port and tries to open a socket listener using these.
// if the `Server struct` is already initialized with an IP and Port, 
// this can function CAN receive both args as `nil`
func (s *Server) Serve (ip string, port int) error {
	err := utils.ValidateIP(ip)

	if err != nil {
		return err
	}

	return nil
}