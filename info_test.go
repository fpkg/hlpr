package hlpr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware_ServiceInfo(t *testing.T) {
	serviceName := "TestService"
	serviceVersion := "v1.0.0"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	ts := httptest.NewServer(ServiceInfo(serviceName, serviceVersion)(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("could nod send request to test server: %v", err)
	}
	defer resp.Body.Close()

	if resp.Header.Get("Service-Name") != serviceName {
		t.Errorf("unexpected service name: %s", resp.Header.Get("Service-Name"))
	}

	if resp.Header.Get("Service-Version") != serviceVersion {
		t.Errorf("unexpected service version: %s", resp.Header.Get("Service-Version"))
	}
}
