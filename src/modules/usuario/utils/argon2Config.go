package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type configArgon2 struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func EncriptarPassword(password *string) (string, error) {
	var config configArgon2 = configArgon2{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	saltos, err := generarRandomBytes(config.saltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(*password), *saltos, config.iterations, config.memory, config.parallelism, config.keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(*saltos)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, config.memory, config.iterations, config.parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func generarRandomBytes(longituSalto uint32) (*[]byte, error) {
	b := make([]byte, longituSalto)

	_, err := rand.Read(b)
	if err != nil {
		return &[]byte{}, err
	}
	return &b, nil
}
