package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	SaltLen uint32
	KeyLen  uint32
}

type HashManager struct {
	Params Argon2Params
}

func NewHashManager(params Argon2Params) *HashManager {
	return &HashManager{
		Params: params,
	}
}

var DefaultHasher = NewHashManager(Argon2Params{
	Time:    1,
	Memory:  64 * 1024,
	Threads: 4,
	SaltLen: 16,
	KeyLen:  32,
})

func (h *HashManager) HashPassword(password string) (string, error) {
	salt := make([]byte, h.Params.SaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		h.Params.Time,
		h.Params.Memory,
		h.Params.Threads,
		h.Params.KeyLen,
	)

	hashStr := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		h.Params.Memory,
		h.Params.Time,
		h.Params.Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return hashStr, nil
}

func slowStringCompare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	result := byte(0)
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}

	return result == 0
}

func (h *HashManager) VerifyPassword(password, hashStr string) bool {
	var version int
	var memory, time uint32
	var threads uint8
	var saltStr, hashHexStr string

	_, err := fmt.Sscanf(
		hashStr,
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		&version,
		&memory,
		&time,
		&threads,
		&saltStr,
		&hashHexStr,
	)
	if err != nil {
		fmt.Printf("Failed to parse hash: %v\n", err)
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Printf("Failed to decode salt: %v\n", err)
		return false
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(hashHexStr)
	if err != nil {
		fmt.Printf("Failed to decode hash: %v\n", err)
		return false
	}

	computedHash := argon2.IDKey(
		[]byte(password),
		salt,
		time,
		memory,
		threads,
		uint32(len(storedHash)),
	)

	return slowStringCompare(string(computedHash), string(storedHash))
}
