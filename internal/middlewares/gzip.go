package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	ct := w.Header().Get("Content-Type")

	if !strings.Contains(ct, "application/json") && !strings.Contains(ct, "text/html") {
		return w.ResponseWriter.Write(b)
	}

	w.Header().Add("Content-Encoding", "gzip")
	return w.Writer.Write(b)
}

func GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ограничимся только данным вариантом, ибо так прописано в треке
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz := gzip.NewWriter(w)
		defer gz.Close()

		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

type gzipReader struct {
	Reader io.ReadCloser
}

func (r gzipReader) Read(b []byte) (int, error) {
	return r.Reader.Read(b)
}

func (r gzipReader) Close() error {
	return r.Reader.Close()
}

func UnzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r.Body = gzipReader{
				Reader: gz,
			}
		}

		next.ServeHTTP(w, r)
	})
}
