package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// CloseableResponseWriter Flushes internal buffer to the output stream
type CloseableResponseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (zip gzipResponseWriter) Write(data []byte) (int, error) {
	return zip.Writer.Write(data)
}

func (zip gzipResponseWriter) Close() {
	zip.Writer.Close()
}

func (zip gzipResponseWriter) Header() http.Header {
	return zip.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (closable closeableResponseWriter) Close() {

}

func GetResponseWriter(w http.ResponseWriter, req *http.Request) CloseableResponseWriter {
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		grw := gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}
		return grw
	}
	return closeableResponseWriter{ResponseWriter: w}
}
