package set1

import (
	"fmt"
)

func HexToBase64(input []byte) []byte {
	bytes := []byte{}
	encoded := []byte{}
	for i:= 0; i < len(input); i+=2 {
		a, err := byteFromHexChar(input[i])
		if err != nil {
			return []byte{}
		}
		b, err := byteFromHexChar(input[i+1])
		if err != nil {
			return []byte{}
		}
		bytes = append(bytes, (a<<4) | b)
	}
	for _, b := range bytes {

	}
	return encoded
}

func byteFromHexChar(char byte) (byte, error) {
	switch {
	case '0' <= char && char <= '9':
		return char - '0', nil
	case 'a' <= char && char <= 'f':
		return char - 'a' + 10, nil
	}

	return 0, fmt.Errorf("Invalid Hex Char: %v", char)
}