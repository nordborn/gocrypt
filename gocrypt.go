// Package gocrypt is a tiny wrapper that
// provides convenient access
// to the most popular encrypting operations:
// HMAC(SHA) and SHA (sha512 and sha256 are supported)
// with output as HexDigest or Base64 strings
// Example:
//
// 	enc1 := gocrypt.New(msg, secret).HMACSHA(gocrypt.SHA512).HexDigest()
//  enc2 := gocrypt.New(msg, nil).SHA(gocrypt.SHA256).Base64()
//
package gocrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

type alg string

const (
	SHA256 = alg("SHA256")
	SHA512 = alg("SHA512")
)

type in struct {
	msg    []byte
	secret []byte
}

type out struct {
	enc []byte
}

// New creates a new hasher
// If you don't need secret (for further SHA()), set nil
func New(msg, secret []byte) *in {
	hd := in{msg, secret}
	return &hd
}

// HMACSHA encrypts msg using secret with provided algorithm
// gocrypt.SHA256 or gocrypt.SHA512
func (i *in) HMACSHA(a alg) *out {
	var fn func() hash.Hash
	switch a {
	case SHA256:
		fn = sha256.New
	case SHA512:
		fn = sha512.New
	}
	hmc := hmac.New(fn, i.secret)
	hmc.Write(i.msg)
	return &out{hmc.Sum(nil)}
}

// SHA encrypts msg with provided algorithm
// gocrypt.SHA256 or gocrypt.SHA512
func (i *in) SHA(a alg) *out {
	var s []byte
	switch a {
	case SHA256:
		arr := sha256.Sum256(i.msg)
		s = arr[:]
	case SHA512:
		arr := sha512.Sum512(i.msg)
		s = arr[:]
	}
	return &out{s}
}

// Base64 returns base64-encoded string of the hash
// from the step HMACSHA() or SHA()
func (o *out) Base64() string {
	return base64.StdEncoding.EncodeToString(o.enc)
}

// HexDigest returns hax digest string of the hash
// from the step HMACSHA() or SHA()
func (o *out) HexDigest() string {
	return hex.EncodeToString(o.enc)
}
