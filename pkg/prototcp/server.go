package prototcp

import (
	"net"
	"net/url"
	"strconv"
)

// ServerURL returns the URL of a PROTOTCP server.
func ServerURL(host string, address *net.TCPAddr, port int) string {
	hostStr := net.JoinHostPort(host, strconv.Itoa(port))
	if address.Zone != "" {
		hostStr = net.JoinHostPort(host+"%"+address.Zone, strconv.Itoa(port))
	}
	return (&url.URL{Scheme: "rosrpc", Host: hostStr}).String()
}

// Server is a TCPROS server.
type Server struct {
	ln net.Listener
}

// NewServer allocates a Server.
func NewServer(address string) (*Server, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{
		ln: ln,
	}, nil
}

// Close closes the server.
func (s *Server) Close() error {
	return s.ln.Close()
}

// Port returns the server port.
func (s *Server) Port() int {
	return s.ln.Addr().(*net.TCPAddr).Port
}

// Accept accepts clients.
func (s *Server) Accept() (*Conn, error) {
	nconn, err := s.ln.Accept()
	if err != nil {
		return nil, err
	}

	return newConn(nconn), err
}
