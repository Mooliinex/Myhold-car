// middlewares/loggerMiddleware.go

package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request to %s %s took %s", r.Method, r.URL.Path, time.Since(startTime))
	})
}
