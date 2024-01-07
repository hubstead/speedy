package speedy

import (
	"net/http"
)

type Route struct {
	Method string
	Path   string

	MiddlewareBefore []Handler
	MiddlewareAfter  []Handler

	Parent   *Route
	Children []*Route
	Handler  Handler
}

func (r *Route) SetMethod(method string) *Route {
	r.Method = method
	return r
}

func (r *Route) SetPath(path string) *Route {
	r.Path = path
	return r
}

func (r *Route) AddMiddleware(handler Handler) *Route {
	r.MiddlewareBefore = append(r.MiddlewareBefore, handler)
	return r
}

func (r *Route) AddMiddlewareAfter(handler Handler) *Route {
	r.MiddlewareAfter = append(r.MiddlewareAfter, handler)
	return r
}

func (r *Route) SetHandler(handler Handler) *Route {
	r.Handler = handler
	return r
}

// partition takes the current method and handlers and creates children instead
func (r *Route) partition() {
	if r.Handler != nil {
		route := &Route{
			Method:  r.Method,
			Handler: r.Handler,
			Parent:  r,
		}
		r.Children = append(r.Children, route)
		r.Method = ""
		r.Handler = nil
	}
}

func (r *Route) AddMethod(method string, handler Handler) *Route {
	// Is there already a handler bound and this would change the routing?
	if r.Handler != nil {
		if r.Method == method {
			r.Handler = handler
			return r
		}

		r.partition()
	}

	route := &Route{
		Path:    "",
		Method:  method,
		Handler: handler,
		Parent:  r,
	}
	r.Children = append(r.Children, route)

	return r
}

func (r *Route) Get(handler Handler) *Route {
	return r.AddMethod(http.MethodGet, handler)
}

func (r *Route) Post(handler Handler) *Route {
	return r.AddMethod(http.MethodGet, handler)
}

func (r *Route) Put(handler Handler) *Route {
	return r.AddMethod(http.MethodGet, handler)
}

func (r *Route) Delete(handler Handler) *Route {
	return r.AddMethod(http.MethodGet, handler)
}

func (r *Route) AddChild(path string, handlers ...Handler) *Route {
	route := &Route{
		Path:   path,
		Parent: r,
	}

	if len(handlers) > 1 {
		route.MiddlewareBefore = handlers[:len(handlers)-1]
		route.Handler = handlers[len(handlers)-1]
	} else if len(handlers) == 1 {
		route.Handler = handlers[0]
	}

	r.Children = append(r.Children, route)

	return route
}

func (r Route) AbsPath() string {
	if r.Parent != nil {
		if r.Path != "" {
			return r.Parent.AbsPath() + "/" + r.Path
		}

		return r.Parent.AbsPath()
	}

	return r.Path
}

func (r *Route) compile(mux *Muxer) {

}

func (m *Muxer) Add(path string, handlers ...Handler) *Route {
	route := &Route{
		Path: path,
	}

	if len(handlers) > 1 {
		route.MiddlewareBefore = handlers[:len(handlers)-1]
		route.Handler = handlers[len(handlers)-1]
	} else if len(handlers) == 1 {
		route.Handler = handlers[0]
	}

	m.routes = append(m.routes, route)

	return route
}
