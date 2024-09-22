package apiserver

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	code int
}

func (w *ResponseWriter) WriteHeader(statuscode int) {
	w.code = statuscode
	w.ResponseWriter.WriteHeader(statuscode)
}
