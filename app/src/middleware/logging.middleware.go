package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	Reset, Green, Yellow, Cyan := "\033[0m", "\033[32m", "\033[33m", "\033[36m"

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", Green+req.Method+Reset, Cyan+req.RequestURI+Reset, Yellow+time.Since(start).String()+Reset)
	})
}
