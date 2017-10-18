package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
	"time"
	"twilight/deciphers"
	"twilight/encryption"
)

func main() {
	fmt.Println("Booted go-cipher!")

	startedAt := time.Now() // Denotes the process start time
	plainText := "suhail"

	key := encryption.GenerateCipher(128)
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	encrypted := encryption.Encrypt(key, plainText, nonce) // Returns slice of bytes

	status := bootDecipher(plainText, encrypted, nonce) // Boots up the decipher system
	if status {
		timeLog(startedAt) // Prints the time log on console
	} else {
		fmt.Println("Deciphers gave up!")
	}
}

func bootDecipher(plainText string, encrypted []byte, nonce []byte) bool {

	return deciphers.Boot(
		plainText,
		encrypted,
		nonce,
		128,
		2,
	)
}

func timeLog(startedAt time.Time) {
	elapsed := time.Now().Sub(startedAt)
	// @link: https://golang.org/pkg/strconv/#example_FormatFloat
	fmt.Println("Hours elapsed: " + strconv.FormatFloat(elapsed.Hours(), 'g', -1, 64))
	fmt.Println("Minutes elapsed: " + strconv.FormatFloat(elapsed.Minutes(), 'g', -1, 64))
	fmt.Println("Seconds elapsed: " + strconv.FormatFloat(elapsed.Seconds(), 'g', -1, 64))
}
