package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// Gzip Compression Writer, used as a replacement param to handler functions that accept http.ResponseWriter
type gzipResponseWriter struct {

	// extends http.ResponseWriter, i.e. hence can be used in its replacement as paramter to functions
	// this attribute is itself the actual response writer of the corresponding request
	http.ResponseWriter
	
	// writing to gzipResponseWriter must write to this 
	// gzip writer
	gzipWriter *gzip.Writer
}

func newgzipResponseWriter(w http.ResponseWriter) (*gzipResponseWriter) {
	return &gzipResponseWriter{ResponseWriter: w, gzipWriter: gzip.NewWriter(w)}
}

func (gzw *gzipResponseWriter) destroygzipResponseWriter() {
	gzw.gzipWriter.Close()
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.gzipWriter.Write(b)
}

func GzipCompressor(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")

		gzw := newgzipResponseWriter(w)
		defer gzw.destroygzipResponseWriter()
		handler.ServeHTTP(gzw, r)
	})
}
