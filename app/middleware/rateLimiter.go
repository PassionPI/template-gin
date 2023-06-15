package middleware

import (
	"fmt"
	"net/http"
	"time"

	"app.land.x/pkg/resp"
	"github.com/gin-gonic/gin"
)

// RateLimiter is a middleware that limits the number of requests per client.
func (m *Middleware) RateLimiter(maxRequests int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := m.core.Rds.Client
		ctx := c.Request.Context()
		key := fmt.Sprintf("%s:%s:%s", c.Request.URL.Path, "middleware.RateLimiter", c.ClientIP())

		count, err := client.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if count == 1 {
			client.Expire(ctx, key, duration)
		}

		if count > int64(maxRequests) {
			// 	"code": http.StatusTooManyRequests,
			resp.Err(c, "请求过多，请稍后重试")
			c.Abort()
			return
		}
		c.Next()
	}
}
