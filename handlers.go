package speedy

import "net/http"

// Handler is the type of function that can respond to a HTTP request. It is
// provided a combined struct of [Conn] which contains the request and response
// objects.
type Handler func(conn *Conn) error

// WrapHandler takes a [speedy.Handler] and converts it for use with [http.HandlerFunc].
func WrapHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(NewConn(w, r))
	}
}

// Handle binds an endpoint handler for a pattern path with no method limitations.
//
// Note that this will be called last in the event that another method specific
// request is present as that handler is considered "more specific".
func (m *Muxer) Handle(path string, handler Handler) {

}

// Get binds an endpoint handler for a pattern path limited to a GET method.
func (m *Muxer) Get(path string, handler Handler) {

}

// Post binds an endpoint handler for a pattern path limited to a POST method.
func (m *Muxer) Post(pattern string, handler Handler) {

}

// Put binds an endpoint handler for a pattern path limited to a PUT method.
func (m *Muxer) Put(pattern string, handler Handler) {

}

// Delete binds an endpoint handler for a pattern path limited to a DELETE method.
func (m *Muxer) Delete(pattern string, handler Handler) {

}
