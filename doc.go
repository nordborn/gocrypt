/*
Package gocrypt is a tiny wrapper that
provides convenient access
to the most popular encrypting operations:
HMAC(SHA) and SHA (sha512 and sha256 are supported)
with output as HexDigest or Base64 strings

Example:
	enc1 := gocrypt.New(msg, secret).HMACSHA(gocrypt.SHA512).HexDigest()
	enc2 := gocrypt.New(msg, nil).SHA(gocrypt.SHA256).Base64()

Call format:
	New(msg, secretOptional).Hasher(Algorithm).Output()

Hashers:
	HMACSHA()
	SHA()

Algorithms:
	gocrypt.SHA256
	gocrypt.SHA512

Output:
	HexDigest()
	Base64()
*/
package gocrypt
