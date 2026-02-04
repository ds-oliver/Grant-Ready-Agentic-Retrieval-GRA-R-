package lock

import (
	"context"
	"errors"
	"sync"
	"time"
)

// LockManager defines a distributed lock manager interface.
type LockManager interface {
	Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error)
	Release(ctx context.Context, key string) error
}

// MockLockManager simulates distributed locks in memory.
type MockLockManager struct {
	mu    sync.Mutex
	locks map[string]time.Time
}

// NewMockLockManager initializes a mock lock manager.
func NewMockLockManager() *MockLockManager {
	return &MockLockManager{locks: make(map[string]time.Time)}
}

// Acquire attempts to acquire a lock for the given key and ttl.
func (m *MockLockManager) Acquire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	_ = ctx
	if ttl <= 0 {
		return false, errors.New("ttl must be positive")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	expiration, exists := m.locks[key]
	if exists && time.Now().Before(expiration) {
		return false, nil
	}

	m.locks[key] = time.Now().Add(ttl)
	return true, nil
}

// Release removes a lock for the given key.
func (m *MockLockManager) Release(ctx context.Context, key string) error {
	_ = ctx
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.locks, key)
	return nil
}
