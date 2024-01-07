package speedy

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Server holds the main data for starting up a web HTTP server.
type Server struct {
	host string
	port uint

	tlsCertPath string
	tlsKeyPath  string

	muxer *Muxer

	mux  *http.ServeMux
	srv1 *http.Server
	srv2 *http2.Server
}

// Address returns the joined host:port address
func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

// ServeMux returns the underlying [http.ServeMux] pointer.
func (s *Server) ServeMux() *http.ServeMux {
	return s.mux
}

// HTTPServer returns the underlying [http.Server] pointer.
func (s *Server) HTTPServer() *http.Server {
	return s.srv1
}

// HTTP2Server returns the underlying [http2.Server] pointer. This object is only
// used in HTTP2 clear-text mode and not during TLS operation.
func (s *Server) HTTP2Server() *http2.Server {
	return s.srv2
}

// UseTLS applies the certificate and key filepaths to the server to allow for
// SSL/TLS on the server. It performs a quick [os.Stat] check to ensure validity
// here so that you don't have to wait for the actual listen to fail.
func (s *Server) UseTLS(certPath, keyPath string) error {
	s.tlsCertPath = certPath
	s.tlsKeyPath = keyPath

	if _, err := os.Stat(s.tlsCertPath); err != nil {
		return err
	}

	if _, err := os.Stat(s.tlsKeyPath); err != nil {
		return err
	}

	return nil
}

// BuildMux assembles the final runtime [http.ServeMux] that will handle the
// requests setup for binding. This is called automatically by the
// [ListenAndServe] method, but is provided in case you need to rebuild it.
func (s *Server) BuildMux() {
	s.mux = s.muxer.BuildServeMux()
}

// ListenAndServe opens and starts listening on the server address for incoming
// requests. If there is a TLS certificate and key filepath set, it will use TLS.
// Otherwise it will use basic HTTP.
func (s *Server) ListenAndServe() error {
	// Rebuild the serve mux based on the templates we provided.
	s.BuildMux()

	if s.tlsCertPath != "" && s.tlsKeyPath != "" {
		s.srv1.Handler = s.mux
		return s.srv1.ListenAndServeTLS(s.tlsCertPath, s.tlsKeyPath)
	}

	// Use x/net/http2/h2c in order to use HTTP2 over clear text
	s.srv1.Handler = h2c.NewHandler(s.mux, s.srv2)

	return s.srv1.ListenAndServe()
}

// NewServer constructs a new [Server] pointer for use and sets the host and
// port parameters.
func NewServer(host string, port uint) *Server {
	return &Server{
		host: host,
		port: port,

		muxer: NewMuxer(),

		mux:  nil,
		srv1: &http.Server{},
		srv2: &http2.Server{},
	}
}
