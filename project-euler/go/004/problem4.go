package main

import "fmt"
import "strconv"

// https://github.com/daviskoh/go-code-challenges/tree/master/reverse_string
func ReverseString(s string) string {
	if len(s) < 2 {
		return s
	}

	return s[len(s)-1:len(s)] + ReverseString(s[0:len(s)-1])
}

func isPalendrome(n int) bool {
	testString := strconv.Itoa(n)
	return testString == ReverseString(testString)
}

func problem4() int {
	max := 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			result := i * j
			if isPalendrome(i*j) && result > max {
				max = result
			}
		}
	}

	return max
}

func main() {
	result := problem4()
	fmt.Println("Answer:", result)
}
