package middleware

import (
	"fmt"
	"net/http"
	"time"

	"app_ink/pkg/util"

	"github.com/gin-gonic/gin"
)

// RateLimiter is a middleware that limits the number of requests per client.
func (m *Middleware) RateLimiter(maxRequests int64, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := m.core.Dep.Rds.Client
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

		if count > maxRequests {
			// "code": http.StatusTooManyRequests,
			util.Bad(c, "请求过多，请稍后重试")
			return
		}
		c.Next()
	}
}
