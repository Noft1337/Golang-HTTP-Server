package server

import (
	"fmt"
	"net"

	"xhttp/http"
	"xhttp/logging"
	"xhttp/utils"
)

const (
	CHANNEL_BUFFER_SIZE = 65536
)

var log *logging.Logger

func init() {
	log = logging.NewLogger(0, logging.LOG_NAME)
}

type Server struct {
	Address		string
	Port		int
	AddrFull	string
	Listener 	net.Listener
	Connections []net.Conn
	chCnt		int
	ch 			chan string
}

func NewHTTPServer(addr string, port int) Server {
	return Server {}
}

func (s *Server) serveConnection (conn net.Conn) {
	var data []byte = make([]byte, CHANNEL_BUFFER_SIZE)
	for {
		b, _ := conn.Read(data)
		if b > 0 {
			http.HandleRequest(data)
		}
	}
}

func (s *Server) handleNewConnection (conn net.Conn) {
	s.Connections = append(s.Connections, conn)

	log.Info("Received a connection from %v", conn.RemoteAddr())

	go s.serveConnection(conn)
}

func (s *Server) listenForConnections () {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Err("Error while listening to new connections: %v", err)
		}
		
		go s.handleNewConnection(conn)
	}
}

func (s *Server) Close () {
	utils.CloseConnections(s.Connections)
	s.Listener.Close()
}

// This function takes the ip and port and tries to open a socket listener using these.
// if the `Server struct` is already initialized with an IP and Port, 
// this can function CAN receive both args as "" and `0`
func (s *Server) Serve (ip string, port int) error {
	if ip == "" {
		ip = "0.0.0.0"
	}
	if port == 0 {
		port = utils.RandomizePort()
	}
	
	err := utils.ValidateIP(ip)
	if err != nil {
		return err
	}
	
	err = utils.ValidatePort(port)
	if err != nil {
		return err
	}

	s.Port = port
	s.Address = ip
	s.AddrFull = fmt.Sprintf("%s:%d", s.Address, s.Port)

	ln, err := net.Listen("tcp", s.AddrFull)
	if err != nil {
		log.Err("Can't Listen on %s", s.AddrFull)
		return err
	}

	s.Listener = ln
	go s.listenForConnections()

	return nil
}