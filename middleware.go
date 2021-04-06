package hlpr

import "net/http"

// Use wraps the handler with middleware
func Use(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	if len(mws) == 0 {
		return h
	}

	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}

	return h
}
