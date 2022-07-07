package limiter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	l := &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimiter{
		Limiter: l,
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	//uri := c.Request.RequestURI
	//fmt.Println(uri)
	//index := strings.Index(uri, "2")
	//if index == -1 {
	//	return uri
	//}
	//
	//return uri[:index]
	return "logistics"
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	fmt.Println("获取令牌")
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	fmt.Println("add令牌")
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}

	return l
}
