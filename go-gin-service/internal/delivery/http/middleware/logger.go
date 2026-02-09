package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type timingWriter struct {
	gin.ResponseWriter
	start time.Time
}

func (w timingWriter) WriteHeader(code int) {
	duration := time.Since(w.start)
	w.ResponseWriter.Header().Set("X-Process-Time", duration.String())
	w.ResponseWriter.WriteHeader(code)
}

/*  MIDDLEWARE  */

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		userAgent := c.GetHeader("User-Agent")
		log.Printf("[LOG] Request received from: %s", userAgent)

		c.Writer = timingWriter{
			ResponseWriter: c.Writer,
			start:          start,
		}

		c.Next()
	}
}
