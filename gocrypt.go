package gocrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

// Specific vars SHA256 and SHA512
// will be used as args for HMACSHA() or SHA()
type alg struct {
	v string
}

var (
	// Use it as the arg for HMACSHA() or SHA()
	SHA256 = &alg{"SHA256"}
	// Use is as the arg for HMACSHA() or SHA()
	SHA512 = &alg{"SHA512"}
)

// In represents incoming data
// and provides corresponding methods
type In struct {
	msg    []byte
	secret []byte
}

// Out represents output data and provides
// corresponding methods
type Out struct {
	enc []byte
}

// New creates a new hasher.
// If you don't need secret (for further SHA()), pass nil
func New(msg, secret []byte) *In {
	hd := In{msg, secret}
	return &hd
}

// NewFromStr accepts strings instead of
// slice of bytes and then calls New()
func NewFromStr(msg, secret string) *In {
	return New([]byte(msg), []byte(secret))
}

// HMACSHA encrypts msg using secret with
// provided algorithm name gocrypt.SHA256 or gocrypt.SHA512
func (i *In) HMACSHA(a *alg) *Out {
	var fn func() hash.Hash
	switch a {
	case SHA256:
		fn = sha256.New
	case SHA512:
		fn = sha512.New
	}
	hmc := hmac.New(fn, i.secret)
	hmc.Write(i.msg)
	return &Out{hmc.Sum(nil)}
}

// SHA encrypts msg with
// provided algorithm name gocrypt.SHA256 or gocrypt.SHA512
func (i *In) SHA(a *alg) *Out {
	var s []byte
	switch a {
	case SHA256:
		arr := sha256.Sum256(i.msg)
		s = arr[:]
	case SHA512:
		arr := sha512.Sum512(i.msg)
		s = arr[:]
	}
	return &Out{s}
}

// Base64 returns base64-encoded string of the hash
// from the step HMACSHA() or SHA()
func (o *Out) Base64() string {
	return base64.StdEncoding.EncodeToString(o.enc)
}

// HexDigest returns hex digest string of the hash
// from the step HMACSHA() or SHA()
func (o *Out) HexDigest() string {
	return hex.EncodeToString(o.enc)
}
