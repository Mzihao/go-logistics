package middleware

import (
	"github.com/gin-gonic/gin"
	"go-logistics/utils/errmsg"
	"go-logistics/utils/limiter"
	"net/http"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				code := errmsg.ErrorTooManyRequests
				c.JSON(http.StatusOK, gin.H{
					"status":  code,
					"message": errmsg.GetErrMsg(code),
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
