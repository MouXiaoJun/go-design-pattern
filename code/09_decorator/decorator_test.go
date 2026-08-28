package _9_decorator

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func do(h http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestAuthReject(t *testing.T) {
	h := Chain(okHandler(), Auth)
	rec := do(h, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("got %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestAuthPass(t *testing.T) {
	h := Chain(okHandler(), Auth)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "token")
	if rec := do(h, req); rec.Code != http.StatusOK {
		t.Fatalf("got %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(1, time.Minute)
	h := rl.Handler(okHandler())
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "1.2.3.4:0"

	if rec := do(h, req); rec.Code != http.StatusOK {
		t.Fatalf("first request got %d, want 200", rec.Code)
	}
	if rec := do(h, req); rec.Code != http.StatusTooManyRequests {
		t.Fatalf("second request got %d, want 429", rec.Code)
	}
}
