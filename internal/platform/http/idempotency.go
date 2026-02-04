package http

import (
	"context"
	"net/http"
)

// IdempotencyStore represents a mockable data access layer for idempotency keys.
type IdempotencyStore interface {
	Exists(ctx context.Context, requestID string) (bool, error)
}

// MockIdempotencyStore is an in-memory mock implementation of IdempotencyStore.
type MockIdempotencyStore struct {
	keys map[string]struct{}
}

// NewMockIdempotencyStore initializes a mock store with optional seed keys.
func NewMockIdempotencyStore(seedKeys ...string) *MockIdempotencyStore {
	keys := make(map[string]struct{}, len(seedKeys))
	for _, key := range seedKeys {
		keys[key] = struct{}{}
	}
	return &MockIdempotencyStore{keys: keys}
}

// Exists reports whether the request ID already exists in the mock store.
func (m *MockIdempotencyStore) Exists(ctx context.Context, requestID string) (bool, error) {
	_ = ctx
	_, ok := m.keys[requestID]
	return ok, nil
}

// IdempotencyMiddleware checks for X-Request-ID headers and returns a cached response if found.
func IdempotencyMiddleware(store IdempotencyStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")
			if requestID != "" {
				exists, err := store.Exists(r.Context(), requestID)
				if err != nil {
					http.Error(w, "idempotency check failed", http.StatusInternalServerError)
					return
				}
				if exists {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"status":"cached","request_id":"` + requestID + `"}`))
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
