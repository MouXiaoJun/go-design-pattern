package _9_decorator

import (
	"log"
	"net/http"
	"time"
)

// Middleware 是 http.Handler 的装饰器抽象：包装一个 Handler，返回一个新的 Handler。
type Middleware func(http.Handler) http.Handler

// Logging 日志装饰器：记录请求与耗时，行为上仍委托给内层 Handler。
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s cost=%s", r.Method, r.URL.Path, time.Since(start))
	})
}

// Auth 鉴权装饰器：没有 token 直接拒绝，否则放行给内层 Handler。
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Chain 按传入顺序逐层叠加装饰器，返回最外层的 Handler。
func Chain(h http.Handler, mw ...Middleware) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return h
}
