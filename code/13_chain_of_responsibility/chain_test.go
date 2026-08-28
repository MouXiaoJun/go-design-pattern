package _13_chain_of_responsibility

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func finalHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func TestChainShortCircuit(t *testing.T) {
	chain := New(finalHandler, &AuthHandler{}, NewRateLimitHandler(2))
	h := chain.Build()

	// 无 token：认证节点短路，不会走到最终处理器。
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", rec.Code)
	}

	// 带 token：通过整条链，最终处理器返回 ok。
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Token", "abc")
	rec = httptest.NewRecorder()
	h(rec, req)
	if rec.Code != http.StatusOK || rec.Body.String() != "ok" {
		t.Fatalf("want 200 ok, got %d %q", rec.Code, rec.Body.String())
	}
}
