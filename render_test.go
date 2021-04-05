package hlpr

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_JSON(t *testing.T) {
	data := struct {
		Foo string `json:"foo"`
	}{
		Foo: "bar",
	}

	wantBody := []byte(`{"foo":"bar"}`)
	wantStatus := http.StatusOK

	handler := func(w http.ResponseWriter, _ *http.Request) {
		RenderJSON(w, data)
	}

	w := httptest.NewRecorder()

	handler(w, nil)

	if w.Result().StatusCode != wantStatus {
		t.Errorf("unexpected status code - want: %d; got: %d", wantStatus, w.Result().StatusCode)
	}

	gotBody := w.Body.Bytes()[:w.Body.Len()-1] // skip new line symbol \n
	if !bytes.Equal(gotBody, wantBody) {
		t.Errorf("unexpected result\nwant: % X\ngot: % X", wantBody, gotBody)
	}
}
