package set1

import (
	"testing"
	. "github.com/onsi/gomega"
)

func TestHexToBase64(t *testing.T) {
	RegisterTestingT(t)
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	result := HexToBase64(input)
	Expect(result).To(Equal([]byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")))
}