package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"twilight/encryption"
)

func main() {
	fmt.Println("Booted go-cipher!")
	key := encryption.GenerateCipher(128)
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	encrypted := encryption.Encrypt(key, "suhail", nonce)
	fmt.Printf("%x\n", encrypted)

	decrypted := encryption.Decrypt(key, encrypted, nonce)
	fmt.Printf("%s\n", decrypted)
}
