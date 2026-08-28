package _13_chain_of_responsibility

import "net/http"

// RateLimitHandler 用信号量限制并发请求数，超过则短路。
type RateLimitHandler struct {
	sem chan struct{}
}

func NewRateLimitHandler(n int) *RateLimitHandler {
	return &RateLimitHandler{sem: make(chan struct{}, n)}
}

func (h *RateLimitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	select {
	case h.sem <- struct{}{}:
		defer func() { <-h.sem }()
		next(w, r)
	default:
		http.Error(w, "too many requests", http.StatusTooManyRequests)
	}
}
