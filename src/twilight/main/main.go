package main

import (
	"fmt"
	"twilight/encryption"
)

func main() {
	fmt.Println("Booted go-cipher!")
	fmt.Println(encryption.GenerateCipher(128))
}
