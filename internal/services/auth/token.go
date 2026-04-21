package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type Token struct {
	Value     string
	UserID    int
	CreatedAt time.Time
	ExpiresAt time.Time
	LastUsed  time.Time
}

type TokenManager struct {
	tokens map[string]*Token
	mu     sync.RWMutex
	ttl    time.Duration // Time to live
}

func NewTokenManager(ttl time.Duration) *TokenManager {
	tm := &TokenManager{
		tokens: make(map[string]*Token),
		ttl:    ttl,
	}

	go tm.cleanupExpiredTokens()

	return tm
}

func (tm *TokenManager) GenerateToken(userID int) (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random token: %w", err)
	}

	tokenValue := hex.EncodeToString(tokenBytes)

	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.tokens[tokenValue] = &Token{
		Value:     tokenValue,
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(tm.ttl),
		LastUsed:  time.Now(),
	}

	return tokenValue, nil
}

func (tm *TokenManager) VerifyToken(tokenValue string) (int, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	token, exists := tm.tokens[tokenValue]
	if !exists {
		return 0, fmt.Errorf("token not found")
	}

	if time.Now().After(token.ExpiresAt) {
		return 0, fmt.Errorf("token expired")
	}

	return token.UserID, nil
}

func (tm *TokenManager) VerifyAndRefresh(tokenValue string) (int, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	token, exists := tm.tokens[tokenValue]
	if !exists {
		return 0, fmt.Errorf("token not found")
	}

	if time.Now().After(token.ExpiresAt) {
		return 0, fmt.Errorf("token expired")
	}

	token.LastUsed = time.Now()
	token.ExpiresAt = time.Now().Add(tm.ttl)

	return token.UserID, nil
}

func (tm *TokenManager) RevokeToken(tokenValue string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, exists := tm.tokens[tokenValue]; !exists {
		return fmt.Errorf("token not found")
	}

	delete(tm.tokens, tokenValue)
	return nil
}

func (tm *TokenManager) cleanupExpiredTokens() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		tm.mu.Lock()

		now := time.Now()
		for tokenValue, token := range tm.tokens {
			if now.After(token.ExpiresAt) {
				delete(tm.tokens, tokenValue)
			}
		}

		tm.mu.Unlock()
	}
}

func (tm *TokenManager) GetTokenInfo(tokenValue string) (*Token, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	token, exists := tm.tokens[tokenValue]
	if !exists {
		return nil, fmt.Errorf("token not found")
	}

	return token, nil
}

func (tm *TokenManager) ActiveTokenCount() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	return len(tm.tokens)
}
