package deciphers

import (
	"fmt"
	"twilight/encryption"
)

type kgenerator struct {
	keyFound bool
}

//Boot ...
func Boot(plainText string, encrypted []byte, nonce []byte, keySize int, threadPoolCount int) bool {
	if len(plainText) == 0 || len(encrypted) == 0 {
		panic("Plain text or encrypted string cannot be empty for the decipher")
	}

	if threadPoolCount == 0 {
		threadPoolCount = 1000 // Thread pool count cannot be zero
		// default to 1000
	}

	status := decipher(plainText, encrypted, nonce, keySize, threadPoolCount)
	if <-status {
		return true
	}
	return false
}

func decipher(plainText string, encrypted []byte, nonce []byte, keySize int, threadPoolCount int) <-chan bool {
	status := make(chan bool)
	go func() {
		foundKey := false
		giveUp := false
		for i := 0; i < threadPoolCount; i++ {
			go func(status chan bool) { // Decipher is one go-routine
				// Multiple deciphers form multiple go-routines
				kgen := kgenerator{keyFound: false}
				keyChannel := keyGenerator(keySize, &kgen)

				for key := range keyChannel {
					if foundKey || giveUp {
						kgen.keyFound = true // This will stop the key generator
						// for all the go-routines

						// Not closing status channel as receiver is not looking
						// for a close but a boolean signal
						// status channel will be garbage collected
						// Note: that it is only necessary to close a channel if the receiver is
						// looking for a close.  Closing the channel is a control signal on the
						// channel indicating that no more data follows.
					}
					decrypted := encryption.Decrypt(key, encrypted, nonce)
					if len(decrypted) > 0 {
						fmt.Println("Encryption key found: " + key)
						status <- true
						foundKey = true
					}
				}
			}(status)
		}

	}()
	return status
}

// Generates and writes the key on a channel
// stream. Generates the key until signalled by
// and external action on kgenerator struct
func keyGenerator(keySize int, kgen *kgenerator) chan string {
	keyCh := make(chan string)
	go func() {
		for !kgen.keyFound { // Controls the loop for key generation
			keyCh <- encryption.GenerateCipher(keySize)
		}

		if kgen.keyFound {
			close(keyCh) // Close the channel if the key was found
		}
	}()
	return keyCh
}
