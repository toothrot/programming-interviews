package main

import "fmt"
import "math/big"

func problem() int {
	count := 0
	for i := 0; count < 10001; i++ {
		b := big.NewInt(int64(i))
		if b.ProbablyPrime(10) {
			count++
			if count == 10001 {
				return i
			}
		}
	}
	return 0
}

func main() {
	result := problem()
	fmt.Println("Answer:", result)
}
