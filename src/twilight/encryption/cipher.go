package encryption

import (
	"math/rand"
	"strconv"
	"time"
)

type cryptogram struct {
	characterList  []string // Slice of characters used to create a cipher
	allowedBitSize [3]int   // Allowed bit size list
}

// GenerateCipher ...
// Generates a random cipher for bit-size
// passed as the argument.
func GenerateCipher(size int) string {
	_cipher := cipherPrototype()

	_allowedBitSize := _cipher.allowedBitSize
	_bitSizeExists := false

	for i := 0; i < len(_allowedBitSize); i++ {
		if size == _allowedBitSize[i] {
			// Check if byte-size passed exists
			_bitSizeExists = true
		}
	}

	if !_bitSizeExists {
		// Throw an error, that byte-size not allowed
		panic("\n\tBit size of " + strconv.Itoa(size) + " not allowed")
	}

	cipherStr := ""
	_charList := _cipher.characterList

	_lenCharList := len(_charList)
	_sizeInBytes := size / 8
	for i := 0; i < _sizeInBytes; i++ {
		// link: https://golang.org/pkg/math/rand/#Rand.Int
		rand.Seed(time.Now().UnixNano())                // Generate a random seed
		cipherStr += _charList[rand.Intn(_lenCharList)] // Append a random character to the cipher string
	}

	return cipherStr
}

// Returns the initialized cipher prototype
func cipherPrototype() cryptogram {
	_charList := getCharacterList()
	_allowedBitSize := getAllowedBitSize()

	_cipher := cryptogram{characterList: _charList, allowedBitSize: _allowedBitSize}

	return _cipher
}

// Generates the character list within
// ascii range 33-126
func getCharacterList() []string {
	var _charList []string // Slice of character list
	for i := 33; i <= 126; i++ {
		_charList = append(_charList, string(i))
	}

	return _charList
}

// Returns the list of allowed bit size
func getAllowedBitSize() [3]int {
	return [3]int{128, 192, 256}
}
