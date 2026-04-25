package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

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

func (h *HashManager) CompareHashes(hashedPassword, storedHash string) bool {
	return subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(storedHash)) == 1
}

func (h *HashManager) IsHash(hash string) error {
	_, err := fmt.Sscanf(
		hash,
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
	)
	return err
}

func (h *HashManager) VerifyPassword(password, hashStr string) bool {
	parts := strings.Split(hashStr, "$")
	if len(parts) != 6 {
		log.Printf("invalid hash format")
		return false
	}

	// parts:
	// [0] "" (empty)
	// [1] "argon2id"
	// [2] "v=<version>"
	// [3] "m=<memory>,t=<time>,p=<threads>"
	// [4] salt (base64)
	// [5] hash (base64)

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		log.Printf("failed to parse version: %v", err)
		return false
	}

	var memory, time uint32
	var threads uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		log.Printf("failed to parse params: %v", err)
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		log.Printf("failed to decode salt: %v", err)
		return false
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		log.Printf("failed to decode hash: %v", err)
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

	return subtle.ConstantTimeCompare(computedHash, storedHash) == 1
}
