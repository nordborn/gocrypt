[![Build Status](https://travis-ci.org/nordborn/gocrypt.svg?branch=master)](https://travis-ci.org/nordborn/gocrypt)
[![Code Coverage](https://codecov.io/gh/nordborn/gocrypt/branch/master/graph/badge.svg)](https://codecov.io/gh/nordborn/gocrypt/branch/master/graph/badge.svg)


**Gocrypt is a tiny convenient wrapper for SHA and HMAC(SHA) functions (supports SHA512 and SHA256)**

It provides simple sequential API to get Base64() or HexDigest().


# Installation

`$ go get -u github.com/nordborn/gocrypt`

# Usage

```Go
import (
    "fmt"
    "github.com/nordborn/gocrypt"
)

func main() {
    msg := []byte("pair=ltcusd&nonce=152442791667649")
    secret := []byte("0IOP8VD6QM5OYM20XSM85IYOU8UHXP36J2RFSFO265J3")
    signed := gocrypt.New(msg, secret).HMACSHA(gocrypt.SHA256).HexDigest()
    fmt.Println(signed)
    // Output: 92e05c5a3d6f2415412030c0ccae1e5e561a41c0aa75a9877b8bbf0cef49cb83
}
```

# Provided functions

**hashers:**
- `HMACSHA()`
- `SHA()`

**algorithms:**
- `gocrypt.SHA256`
- `gocrypt.SHA512`

**output string as:** 
- `HexDigest()`
- `Base64()`

See more in `gocrypt_test.go`

Enjoy!