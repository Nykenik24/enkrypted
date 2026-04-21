package auth

import "time"

type AuthService struct {
	Hasher   *HashManager
	TokenMan *TokenManager
}

func NewAuthService() *AuthService {
	return &AuthService{
		Hasher:   DefaultHasher,
		TokenMan: NewTokenManager(1 * time.Hour),
	}
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
