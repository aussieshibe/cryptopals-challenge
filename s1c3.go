package main

import (
	"encoding/hex"
	"fmt"
	"unicode"
)

func main () {
	inStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	scores := make([]float32, 256)

	// Attempt to decode the hex string into byte array
	byteArray, err := hex.DecodeString(inStr)
	// If unsuccessful, print the error and return
	if err != nil {
		fmt.Println("error:", err)
		return
	}


	for i := byte(0); i <= 254; i++ {
		scores[i] = scoreText(string(oneCharXOR(i, byteArray)[:]))
	}

	// Create array of keys ordered by the score their output gives
	ordered := orderScores(scores)

	for i := byte(0); i < 10; i++ {
		fmt.Println(string(ordered[i]), ", Score:", scores[ordered[i]])
		fmt.Printf("String: ")
		theString := string(oneCharXOR(ordered[i], byteArray)[:])
		for _, runeValue := range theString {
			if unicode.IsPrint(runeValue) {
				fmt.Printf(string(runeValue))
			}
		}
		fmt.Println("\n")
	}

	fmt.Println("Done.")
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

func oneCharXOR(c byte, byteArray []byte) []byte {
	xordByteArray := make([]byte, len(byteArray))
	for i := 0; i < len(byteArray); i++ {
		xordByteArray[i] = c ^ byteArray[i]
	}
	return xordByteArray
}

/**
 * Accepts a string and returns a score, hopefully reflecting
 * 			how likely it is to be sensible and not random garbage
 *
 * Currently this is being done by calculating the vowel to
 * other letters ratio (the higher the better)
 */
func scoreText(str string) float32{
	vowels := "aeiouAEIOU"
	vowelcount := 0

	for i := 0; i < len(str); i++ {
		if contains(string(str[i]), vowels) {
			vowelcount++
		}
	}
	return float32(vowelcount) / float32(len(str) - vowelcount)
}

/**
 * Returns true if str contains any characters in chars
 * Is case sensitive (i.e. contains("Hello", "h") == false)
 */
func contains(str string, chars string) bool {
	for c := range chars {
		for s := range str {
			if chars[c] == str[s] {
				return true
			}
		}
	}
	return false
}

func orderScores(scores []float32) []byte {
	originalScores := make([]float32, len(scores))
	copy(originalScores, scores)
	ordered := make([]byte, len(scores))
	for i := range scores {
		var maxIndex byte = 0
		for j := range originalScores {
			if originalScores[j] > originalScores[maxIndex] {
				maxIndex = byte(j)
			}
		}
		ordered[i] = maxIndex
		originalScores[maxIndex] = -1
	}
	return ordered;
}
