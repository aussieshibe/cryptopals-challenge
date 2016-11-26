package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main () {
	hexstr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	// Attempt to decode the hex string into a byte array
	byteArray, err := hex.DecodeString(hexstr)
	// If unsuccessful, print the error and return
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// Encode the byte array as a base 64 string and print it
	base64String := base64.StdEncoding.EncodeToString(byteArray)
	fmt.Printf("%q\n", base64String)
}
