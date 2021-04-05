package hlpr

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RenderJSON serialize data to json and write it to http response
func RenderJSON(w http.ResponseWriter, data interface{}) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}
