package main

import (
	"testing"
)

var secret = "0IOP8VD6QM5OYM20XSM85IYOU8UHXP36J2RFSFO265J3"
var msg = "pair=ltcusd&nonce=152442791667649"

func compare(t *testing.T, res, expected string) {
	if res != expected {
		t.Error("RES:", res, "!=", "EXPECTED:", expected)
	}
}

func TestHMACSHA512Base64Encoded(t *testing.T) {
	res := HMACSHA512Base64Encoded(secret, msg)
	expected := "XPMgjBZOWYPG3VIbfLDXBLCwx00syKMXhlDvrMExb8fFrrK6P8lHNAEtvKfDyalGy5UtZJ2Txody1w+SD26Vlw=="
	compare(t, res, expected)
}

func TestHMACSHA512HexDigest(t *testing.T) {
	res := HMACSHA512HexDigest(secret, msg)
	expected := "5cf3208c164e5983c6dd521b7cb0d704b0b0c74d2cc8a3178650efacc1316fc7c5aeb2ba3fc94734012dbca7c3c9a946cb952d649d93c68772d70f920f6e9597"
	compare(t, res, expected)
}

func TestHMACSHA256Base64Encoded(t *testing.T) {
	res := HMACSHA256Base64Encoded(secret, msg)
	expected := "kuBcWj1vJBVBIDDAzK4eXlYaQcCqdamHe4u/DO9Jy4M="
	compare(t, res, expected)
}

func TestHMACSHA256HexDigest(t *testing.T) {
	res := HMACSHA256HexDigest(secret, msg)
	expected := "92e05c5a3d6f2415412030c0ccae1e5e561a41c0aa75a9877b8bbf0cef49cb83"
	compare(t, res, expected)
}

func TestSHA256HexDigest(t *testing.T) {
	res := SHA256HexDigest(msg)
	expected := "d63276a0b807aed965f6e2d1d142b7ce6fa25d9701dcd9aa3387d74b0008c8d5"
	compare(t, res, expected)
}

func TestSHA512HexDigest(t *testing.T) {
	res := SHA512HexDigest(msg)
	expected := "c7e9f2d16e525212d1651cc6eaf1af3d3bb8058694986f5a41b5d6e1a0f92cefa4d3deec4d3e977a16eb258acdf62ad896b178e75812cdb032177478afa4d46b"
	compare(t, res, expected)
}

func TestSHA256Base64Encoded(t *testing.T) {
	res := SHA256Base64Encoded(msg)
	expected := "1jJ2oLgHrtll9uLR0UK3zm+iXZcB3NmqM4fXSwAIyNU="
	compare(t, res, expected)
}

func TestSHA512Base64Encoded(t *testing.T) {
	res := SHA512Base64Encoded(msg)
	expected := "x+ny0W5SUhLRZRzG6vGvPTu4BYaUmG9aQbXW4aD5LO+k097sTT6XehbrJYrN9irYlrF451gSzbAyF3R4r6TUaw=="
	compare(t, res, expected)
}
