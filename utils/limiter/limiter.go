package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(c *gin.Context) string                          // 获取限流器名称
	GetBucket(key string) (*ratelimit.Bucket, bool)     // 获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // 新增令牌桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket // 令牌桶与名称映射关系
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration // 间隔多久放N个令牌
	Capacity     int64         // 令牌桶容量
	Quantum      int64         // 每次达到间隔时间后所放的具体令牌数量
}
