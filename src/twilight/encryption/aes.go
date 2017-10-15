package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

//Encrypt ...
// Does the AES encryption with the key passed
// and returns the encrypted string
// key (string) is the key to encrypt the
// text (string) passed
func Encrypt(key string, text string) []byte {

	if len(text) == 0 {
		panic("Text cannot be empty")
	}

	// Create a cipher-block
	//@see https://golang.org/pkg/crypto/aes/#NewCipher
	//Note: The key argument should be the AES key, either 16, 24, or 32 bytes
	//to select AES-128, AES-192, or AES-256.
	if validateEncryptionKey(key) {
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			panic(err.Error())
		}
		// NewGCM: Block cipher mode of operation (Algorithm that uses block cipher
		// to provide an information service such as confidentiality or authenticity)
		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}
		nonce := make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}
		encrypted := aesgcm.Seal(nil, nonce, []byte(text), nil)
		return encrypted
	}

	panic("Invalid encryption key passed. Check length of the key.")
}

// Validate the size of the encryption key
// Returns true upon successfull validation
// and false upon failure
func validateEncryptionKey(keys string) bool {
	_len := len(keys)
	if _len > 0 && (_len == 16 || _len == 24 || _len == 32) {
		return true
	}
	return false

}
