package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		ua := c.GetHeader("User-Agent")
		fmt.Println("[LOG] Request received from:", ua)
		c.Next()
		duration := time.Since(start)
		c.Writer.Header().Set("X-Process-Time", duration.String())
	}

}
