package main

import (
	"encoding/hex"
	"fmt"
)

func main () {
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"

	// Attempt to decode the hex strings into byte arrays
	byteArray1, err := hex.DecodeString(str1)
	byteArray2, err := hex.DecodeString(str2)
	// If unsuccessful, print the error and return
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// XOR data
	xordByteArray := make([]byte, len(byteArray1))
	xorBytes(xordByteArray, byteArray1, byteArray2) 
	// Encode the byte array as a hex string and print it
	xorHexString := hex.EncodeToString(xordByteArray)
	fmt.Printf("%q\n", xorHexString)
}

/*
 * XOR function largely borrowed from 
 * golang.org/src/crypto/ciper/xor.go
 */
func xorBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}