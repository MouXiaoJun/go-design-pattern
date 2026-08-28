package _13_chain_of_responsibility

import "net/http"

// AuthHandler 校验请求头中的 token，失败则短路，不再交给下一个节点。
type AuthHandler struct{}

func (AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get("X-Token") == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	next(w, r)
}
