package http

// Gateway defines an HTTP gateway for a rating service.
type Gateway struct {
	addr string
}

// New creates a new HTTP gateway for a rating service.
func New(addr string) *Gateway {
	return &Gateway{addr}
}
