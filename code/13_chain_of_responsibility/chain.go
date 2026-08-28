package _13_chain_of_responsibility

import "net/http"

// Handler 责任链上的一个节点：处理请求，或把它交给下一个节点。
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

// HandlerFunc 适配普通函数，使其满足 Handler 接口。
type HandlerFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	f(w, r, next)
}

// Chain 把一串中间件和一个最终处理器串成一条链。
type Chain struct {
	handlers []Handler
	final    http.HandlerFunc
}

func New(final http.HandlerFunc, handlers ...Handler) *Chain {
	return &Chain{handlers: handlers, final: final}
}

// Build 从后往前组装：每个中间件拿到的 next 是「它后面所有节点串起来的结果」。
func (c *Chain) Build() http.HandlerFunc {
	next := c.final
	for i := len(c.handlers) - 1; i >= 0; i-- {
		h := c.handlers[i] // 每次迭代新建变量，闭包各捕获一份
		prev := next       // 快照当前 next，避免闭包捕获循环变量
		next = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r, prev)
		}
	}
	return next
}
