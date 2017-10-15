package main

import (
	"fmt"
	"twilight/encryption"
)

func main() {
	fmt.Println("Booted go-cipher!")
	key := encryption.GenerateCipher(128)
	encrypted := encryption.Encrypt(key, "suhail")
	fmt.Println(string(encrypted))
}
