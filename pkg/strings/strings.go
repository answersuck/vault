package strings

import (
	cryptoRand "crypto/rand"
	mathRand "math/rand"
)

const (
	chars   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// NewUnique generates random string using
// Cryptographically Secure Pseudorandom number
func NewUnique(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := cryptoRand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

// NewRandom generates random URL safe string
func NewRandom(length int) string {
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = chars[mathRand.Intn(len(chars))]
	}

	return string(bytes)
}

// NewSpecialRandom generates random string with special characters
func NewSpecialRandom(length int) string {
	bytes := make([]byte, length)

	c := chars + special
	for i := range bytes {
		bytes[i] = c[mathRand.Intn(len(chars))]
	}

	return string(bytes)
}
