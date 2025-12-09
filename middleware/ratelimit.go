package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"sync"
	"time"
	"video-api/handler"
	"video-api/pkg/errno"
)

// 针对ip进行限流
type IPRateLimiter struct {
	ips  sync.Map
	rate float64
	cap  int64
}

func NewIPRateLimiter(r float64, c int64) *IPRateLimiter {
	return &IPRateLimiter{
		rate: r,
		cap:  c,
	}
}
func (i *IPRateLimiter) GetBucket(ip string) *ratelimit.Bucket {
	bucket, exists := i.ips.Load(ip)
	if exists {
		return bucket.(*ratelimit.Bucket)
	}
	newBucket := ratelimit.NewBucketWithQuantum(time.Second, i.cap, int64(i.rate))
	i.ips.Store(ip, newBucket)
	return newBucket
}
func RateLimitMiddleware() gin.HandlerFunc {
	limiter := NewIPRateLimiter(10, 20)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		bucket := limiter.GetBucket(ip)
		if bucket.TakeAvailable(1) == 0 {
			handler.SendResponse(c, errno.TooManyRequestErr, nil)
			//作用是停止中间件链条，如果没有这行，请求还会继续返回给后面的业务逻辑
			//导致限流失败，依然消耗资源
			c.Abort()
			return
		}
		c.Next()
	}
}
