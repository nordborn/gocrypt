// Package gocrypt is a tiny wrapper that
// provides convenient access
// to the most used encrypting operations
package gocrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

// todo impl interface
// New(msg, secret).HMAC(gocrypt.SHA265).HexDigest()
// New(msg, secret).SHA256().Base64()

func hmacdo(secret, msg string, shaFn func() hash.Hash) []byte {
	h := hmac.New(shaFn, []byte(secret))
	h.Write([]byte(msg))
	return h.Sum(nil)
}

// HMACSHA256Base64Encoded as a string
func HMACSHA256Base64Encoded(secret, msg string) string {
	hd := hmacdo(secret, msg, sha256.New)
	return base64.StdEncoding.EncodeToString(hd)
}

// HMACSHA256HexDigest as a string
func HMACSHA256HexDigest(secret, msg string) string {
	hd := hmacdo(secret, msg, sha256.New)
	return hex.EncodeToString(hd)
}

// HMACSHA512Base64Encoded as a string
func HMACSHA512Base64Encoded(secret, msg string) string {
	hd := hmacdo(secret, msg, sha512.New)
	return base64.StdEncoding.EncodeToString(hd)
}

// HMACSHA512HexDigest as a string
func HMACSHA512HexDigest(secret, msg string) string {
	hd := hmacdo(secret, msg, sha512.New)
	return hex.EncodeToString(hd)
}

// SHA256Base64Encoded as a string
func SHA256Base64Encoded(msg string) string {
	s := sha256.Sum256([]byte(msg))
	return base64.StdEncoding.EncodeToString(s[:])
}

// SHA256HexDigest as a string
func SHA256HexDigest(msg string) string {
	s := sha256.Sum256([]byte(msg))
	return hex.EncodeToString(s[:])
}

// SHA512Base64Encoded as a string
func SHA512Base64Encoded(msg string) string {
	s := sha512.Sum512([]byte(msg))
	return base64.StdEncoding.EncodeToString(s[:])
}

// SHA512HexDigest as a string
func SHA512HexDigest(msg string) string {
	s := sha512.Sum512([]byte(msg))
	return hex.EncodeToString(s[:])
}
