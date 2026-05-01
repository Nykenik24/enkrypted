package auth

import (
	"time"
)

type AuthSession struct {
	AdminHash string
}

type AuthService struct {
	Hasher   *HashManager
	TokenMan *TokenManager
	Session  *AuthSession
}

func NewAuthService(adminPasswd string) (*AuthService, error) {
	serv := &AuthService{
		Hasher:   DefaultHasher,
		TokenMan: NewTokenManager(1 * time.Hour),
	}

	hashed, err := serv.Hasher.HashPassword(adminPasswd)
	if err != nil {
		return nil, err
	}

	serv.Session = &AuthSession{AdminHash: hashed}

	return serv, nil
}

func (a *AuthService) GetAdminHash() string {
	return a.Session.AdminHash
}

type AuthConfig struct {
	Ar2Params Argon2Params
	TokenTTL  time.Duration
}

func NewAdvancedAuthService(config AuthConfig) *AuthService {
	return &AuthService{
		Hasher:   NewHashManager(config.Ar2Params),
		TokenMan: NewTokenManager(config.TokenTTL),
	}
}
