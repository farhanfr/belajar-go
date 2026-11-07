package middlewares

import (
	"log"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
}

func (rec *statusRecorder) WriteHeader(code int) {
	if rec.wroteHeader {
		return
	}

	rec.statusCode = code
	rec.wroteHeader = true
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *statusRecorder) Write(body []byte) (int, error) {
	if !rec.wroteHeader {
		rec.WriteHeader(http.StatusOK)
	}

	return rec.ResponseWriter.Write(body)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		defer func() {
			duration := time.Since(start)

			log.Printf(
				"%s %s | Status: %d | Duration: %s",
				r.Method,
				r.URL.Path,
				recorder.statusCode,
				duration,
			)
		}()

		next.ServeHTTP(recorder, r)
	})
}
