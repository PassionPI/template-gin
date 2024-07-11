package middleware

// import (
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/rs/zerolog/log"
// )

// type LoggerConfig struct{}

// // LoggerWithConfig instance a Logger middleware with config.
// func LoggerWithConfig(_ LoggerConfig) gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		// Start timer
// 		start := time.Now()
// 		path := c.Request.URL.Path
// 		raw := c.Request.URL.RawQuery

// 		// Process request
// 		c.Next()

// 		// Stop timer
// 		end := time.Now()
// 		cost := end.Sub(start)
// 		ip := c.ClientIP()
// 		method := c.Request.Method
// 		status := c.Writer.Status()
// 		errorMessage := c.Errors.ByType(gin.ErrorTypeAny).String()
// 		bodySize := c.Writer.Size()

// 		if raw != "" {
// 			path = path + "?" + raw
// 		}

// 		log.Info()

// 	}
// }
