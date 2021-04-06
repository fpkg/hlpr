package hlpr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware_Use(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/t/foo/bar" {
			t.Errorf("unexpected result: %s", r.URL.Path)
		}
	})

	mw1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path += "/foo"
			next.ServeHTTP(w, r)
		})
	}

	mw2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path += "/bar"
			next.ServeHTTP(w, r)
		})
	}

	ts := httptest.NewServer(Use(handler, mw1, mw2))
	defer ts.Close()

	_, err := http.Get(ts.URL + "/t")
	if err != nil {
		t.Errorf("could not get test http server: %v", err)
	}
}
