package crypto

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const salt_size = 16

type Argon2Config struct {
	HashRaw    []byte
	Salt       []byte
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

func generateCryptographicSalt(saltSize uint32) ([]byte, error) {
	salt := make([]byte, saltSize)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("salt generation failed: %w", err)
	}

	return salt, nil
}

func HashPasswordSecure(password string) (string, error) {
	config := &Argon2Config{
		TimeCost:   2,
		MemoryCost: 64 * 1024,
		Threads:    4,
		KeyLength:  32,
	}

	salt, err := generateCryptographicSalt(salt_size)
	if err != nil {
		return "", fmt.Errorf("password hashish failed: %w", err)
	}
	config.Salt = salt

	config.HashRaw = argon2.IDKey(
		[]byte(password),
		config.Salt,
		config.TimeCost,
		config.MemoryCost,
		config.Threads,
		config.KeyLength,
	)

	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		config.MemoryCost,
		config.TimeCost,
		config.Threads,
		base64.RawStdEncoding.EncodeToString(config.Salt),
		base64.RawStdEncoding.EncodeToString(config.HashRaw),
	)

	return encodedHash, nil
}

func parseArgon2Hash(encodedHash string) (*Argon2Config, error) {
	components := strings.Split(encodedHash, "$")
	if len(components) != 6 {
		return nil, errors.New("invalid hash format structure")
	}

	hashAlgorithm := components[1]
	hashVersion := components[2]
	hashParameters := components[3]
	hashSaltComponent := components[4]
	hashComponent := components[5]

	if !strings.HasPrefix(hashAlgorithm, "argon2id") {
		return nil, errors.New("unsupported algorithm variant")
	}

	var version int
	fmt.Sscanf(hashVersion, "v=%d", &version)

	config := &Argon2Config{}
	fmt.Sscanf(hashParameters, "m=%d,t=%d,p=%d", &config.MemoryCost, &config.TimeCost, &config.Threads)

	salt, err := base64.RawStdEncoding.DecodeString(hashSaltComponent)
	if err != nil {
		return nil, fmt.Errorf("salt decoding failed: %w", err)
	}
	config.Salt = salt

	hash, err := base64.RawStdEncoding.DecodeString(hashComponent)
	if err != nil {
		return nil, fmt.Errorf("hash decoding failed: %w", err)
	}
	config.HashRaw = hash
	config.KeyLength = uint32(len(hash))

	return config, nil
}

func VerifyPasswordSecure(storedHash, providedPassword string) (bool, error) {
	config, err := parseArgon2Hash(storedHash)
	if err != nil {
		return false, fmt.Errorf("hash parsing failed: %w", err)
	}

	computedHash := argon2.IDKey(
		[]byte(providedPassword),
		config.Salt,
		config.TimeCost,
		config.MemoryCost,
		config.Threads,
		config.KeyLength,
	)

	match := subtle.ConstantTimeCompare(config.HashRaw, computedHash) == 1
	return match, nil
}
