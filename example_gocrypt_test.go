package gocrypt_test

import (
	"fmt"
	"github.com/nordborn/gocrypt"
)

// This example demonstrates common usage of gocrypt
func Example_HMACSHA() {
	res := gocrypt.NewFromStr("pair=ltcusd&nonce=152442791667649",
		"0IOP8VD6QM5OYM20XSM85IYOU8UHXP36J2RFSFO265J3").
		HMACSHA(gocrypt.SHA256).
		HexDigest()

	fmt.Println(res)
	// Output: 92e05c5a3d6f2415412030c0ccae1e5e561a41c0aa75a9877b8bbf0cef49cb83
}
