package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

//Encrypt ...
// Does the AES encryption with the key passed
// and returns the encrypted string
// key (string) is the key to encrypt the
// text (string) passed
func Encrypt(key string, text string, nonce []byte) []byte {

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

		encrypted := aesgcm.Seal(nil, nonce, []byte(text), nil)
		return encrypted
	}

	panic("Invalid encryption key passed. Check length of the key.")
}

//Decrypt ...
// Decryptes the slice of bytes existing in encrypted
// form using the key (string) passed
// Returning an empty slice of bytes signifies that decryption
// attempt failed for the key passed
func Decrypt(key string, encrypted []byte, nonce []byte) []byte {
	if len(encrypted) == 0 {
		panic("Slice of encrypted bytes cannot be empty")
	}

	if validateEncryptionKey(key) {
		_encrypted, err := hex.DecodeString(hex.EncodeToString(encrypted))
		_nonce, err := hex.DecodeString(hex.EncodeToString(nonce))

		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			panic(err.Error())
		}

		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}
		plaintext, err := aesgcm.Open(nil, _nonce, _encrypted, nil)
		if err != nil {
			plaintext = make([]byte, 0) // Create an empty slice to return
			logger(key, err.Error())
		} else {
			logger(key, "Successfully decrypted")
		}
		return plaintext
	}
	panic("Invalid encryption key passed. Check length of the key passed")
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

func logger(key string, message string) {
	_message := "Message: " + message
	_message += " ... Key used: " + key

	fmt.Println(_message)
}
