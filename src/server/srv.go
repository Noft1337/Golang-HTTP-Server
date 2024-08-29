package server

import (
	// "xhttp/src/log"
	// "xhttp/src/utils"
)

type Server struct {
	Address		string
	Port		int
}

// This function takes the ip and port and tries to open a socket listener using these.
// if the `Server struct` is already initialized with an IP and Port, 
// this can function CAN receive both args as `nil`
func (s *Server) Serve (ip string, port int) {
}