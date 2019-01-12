package gocrypt_test

import (
	"fmt"
	"github.com/nordborn/gocrypt"
	"testing"
)

const (
	msg    = "pair=ltcusd&nonce=152442791667649"
	secret = "0IOP8VD6QM5OYM20XSM85IYOU8UHXP36J2RFSFO265J3"
)

func compare(t *testing.T, res, expected string) {
	if res != expected {
		t.Error("RES:", res, "!=", "EXPECTED:", expected)
	}
}

var expected = map[string]string{
	"HMACSHA512Base64Encoded": "XPMgjBZOWYPG3VIbfLDXBLCwx00syKMXhlDvrMExb8fFrrK6P8lHNAEtvKfDyalGy5UtZJ2Txody1w+SD26Vlw==",
	"HMACSHA512HexDigest":     "5cf3208c164e5983c6dd521b7cb0d704b0b0c74d2cc8a3178650efacc1316fc7c5aeb2ba3fc94734012dbca7c3c9a946cb952d649d93c68772d70f920f6e9597",
	"HMACSHA256Base64Encoded": "kuBcWj1vJBVBIDDAzK4eXlYaQcCqdamHe4u/DO9Jy4M=",
	"HMACSHA256HexDigest":     "92e05c5a3d6f2415412030c0ccae1e5e561a41c0aa75a9877b8bbf0cef49cb83",
	"SHA256HexDigest":         "d63276a0b807aed965f6e2d1d142b7ce6fa25d9701dcd9aa3387d74b0008c8d5",
	"SHA512HexDigest":         "c7e9f2d16e525212d1651cc6eaf1af3d3bb8058694986f5a41b5d6e1a0f92cefa4d3deec4d3e977a16eb258acdf62ad896b178e75812cdb032177478afa4d46b",
	"SHA256Base64Encoded":     "1jJ2oLgHrtll9uLR0UK3zm+iXZcB3NmqM4fXSwAIyNU=",
	"SHA512Base64Encoded":     "x+ny0W5SUhLRZRzG6vGvPTu4BYaUmG9aQbXW4aD5LO+k097sTT6XehbrJYrN9irYlrF451gSzbAyF3R4r6TUaw==",
}

func TestIn_HMACSHA(t *testing.T) {
	h := gocrypt.NewFromBytes([]byte(msg), []byte(secret))
	compare(t, h.HMACSHA(gocrypt.SHA512).Base64(), expected["HMACSHA512Base64Encoded"])
	compare(t, h.HMACSHA(gocrypt.SHA512).HexDigest(), expected["HMACSHA512HexDigest"])
	compare(t, h.HMACSHA(gocrypt.SHA256).Base64(), expected["HMACSHA256Base64Encoded"])
	compare(t, h.HMACSHA(gocrypt.SHA256).HexDigest(), expected["HMACSHA256HexDigest"])

}

func TestIn_SHA(t *testing.T) {
	h := gocrypt.NewFromBytes([]byte(msg), []byte(secret))
	compare(t, h.SHA(gocrypt.SHA256).HexDigest(), expected["SHA256HexDigest"])
	compare(t, h.SHA(gocrypt.SHA512).HexDigest(), expected["SHA512HexDigest"])
	compare(t, h.SHA(gocrypt.SHA256).Base64(), expected["SHA256Base64Encoded"])
	compare(t, h.SHA(gocrypt.SHA512).Base64(), expected["SHA512Base64Encoded"])
}

func TestNew(t *testing.T) {
	h := gocrypt.New(msg, secret)
	compare(t, h.HMACSHA(gocrypt.SHA512).HexDigest(), expected["HMACSHA512HexDigest"])
}

// This example demonstrates common usage of gocrypt
func Example() {
	res := gocrypt.New("pair=ltcusd&nonce=152442791667649",
		"0IOP8VD6QM5OYM20XSM85IYOU8UHXP36J2RFSFO265J3").
		HMACSHA(gocrypt.SHA256).
		HexDigest()

	fmt.Println(res)

	// Output: 92e05c5a3d6f2415412030c0ccae1e5e561a41c0aa75a9877b8bbf0cef49cb83
}
