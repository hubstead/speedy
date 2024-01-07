package speedy

import "net/http"

// Muxer is an extension on the [http.ServeMux] that handles routes. The final
// [http.ServeMux] can be built during runtime.
type Muxer struct {
	// middlewares holds the root middlewares added with [Use]
	middlewares []Handler

	routes []*Route
}

func (m Muxer) BuildServeMux() *http.ServeMux {
	sm := http.NewServeMux()

	return sm
}

// Use binds middleware to the root [Server] handler for use on all requests.
// These are applied in order of usage.
func (m *Muxer) Use(handler Handler) {

}

func NewMuxer() *Muxer {
	return &Muxer{
		middlewares: make([]Handler, 0),
		routes:      make([]*Route, 0),
	}
}
