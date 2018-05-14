package http

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/Tinee/doit/domain"
)

// Server represents the server
type Server struct {
	ln      net.Listener
	handler http.Handler
	Addr    string
}

// Handler represents the core handler
type Handler struct {
	UserHandler *UserHandler
}

//NewServer creates a pointer to a Server struct
func NewServer(port int, r domain.Client) *Server {
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	userRepo := r.UserRepository()
	return &Server{
		Addr: addr,
		handler: Handler{
			UserHandler: NewCredentialHandler(userRepo),
		},
	}
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/user") {
		h.UserHandler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

//Open tries to open a connection to the given port.
func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	go func() { http.Serve(s.ln, s.handler) }()

	return nil
}

// Close attempts to close the listner on the *Server
func (s *Server) Close() error {
	if s.ln == nil {
		return errors.New("Open a connection before you try to close it")
	}

	s.ln.Close()

	return nil
}
