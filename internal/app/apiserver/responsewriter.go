package apiserver

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	code int
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.code = code
	w.WriteHeader(code)
}
