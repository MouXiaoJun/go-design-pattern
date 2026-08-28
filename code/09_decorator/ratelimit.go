package _9_decorator

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter 是有状态的装饰器：固定窗口内限制每个客户端的请求次数。
type RateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	count   map[string]int
	resetAt map[string]time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		count:   make(map[string]int),
		resetAt: make(map[string]time.Time),
	}
}

// Handler 返回限流装饰器，使用客户端地址作为计数 key。
func (rl *RateLimiter) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RemoteAddr
		now := time.Now()

		rl.mu.Lock()
		defer rl.mu.Unlock()

		if rl.resetAt[key].Before(now) {
			rl.count[key] = 0
			rl.resetAt[key] = now.Add(rl.window)
		}
		if rl.count[key] >= rl.limit {
			http.Error(w, "rate limited", http.StatusTooManyRequests)
			return
		}
		rl.count[key]++
		next.ServeHTTP(w, r)
	})
}
