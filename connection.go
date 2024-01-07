package speedy

import "net/http"

// Conn is used to provide a combined handler experience simplifying the
// process of handling connections
type Conn struct {
}

// NewConn wraps a standard [http.HandlerFunc] parameters into a new connection.
func NewConn(w http.ResponseWriter, r *http.Request) *Conn {
	return &Conn{}
}
