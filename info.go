package hlpr

import "net/http"

// ServiceInfo set service name and version to response header
func ServiceInfo(name, version string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Service-Name", name)
			w.Header().Set("Service-Version", version)
			next.ServeHTTP(w, r)
		})
	}
}
