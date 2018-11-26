[![Build Status](https://travis-ci.org/nordborn/gocrypt.svg?branch=master)](https://travis-ci.org/nordborn/gocrypt)
[![Code Coverage](https://codecov.io/gh/nordborn/gocrypt/branch/master/graph/badge.svg)](https://codecov.io/gh/nordborn/gocrypt/branch/master/graph/badge.svg)


**The most popular encrypt functions for Go (Golang) programming language**

# Installation

`$ go get -u github.com/nordborn/gocrypt`

if you use `go.mod`, add:

`github.com/nordborn/gocrypt v1.0.0`

# Usage

```Go
import (
    "fmt"
    "github.com/nordborn/gocrypt"
)

func main() {
    secret := "0IOP8VD6QM5OYM20XSM85IYOU8UHXP36J2RFSFO265J3"
    msg := "pair=ltcusd&nonce=152442791667649"
    signed := gocrypt.HMACSHA256HexDigest(secret, msg)
    // 92e05c5a3d6f2415412030c0ccae1e5e561a41c0aa75a9877b8bbf0cef49cb83
    fmt.Println(signed)
}
```

# Provided functions

```
func HMACSHA256Base64Encoded(secret, msg string) string

func HMACSHA256HexDigest(secret, msg string) string

func HMACSHA512Base64Encoded(secret, msg string) string

func HMACSHA512HexDigest(secret, msg string) string

func SHA256Base64Encoded(msg string) string

func SHA256HexDigest(msg string) string

func SHA512Base64Encoded(msg string) string

func SHA512HexDigest(msg string) string
```