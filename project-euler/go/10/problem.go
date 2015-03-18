package main

import "fmt"
import "math/big"

func problem() int {
	sum := 0
	for i := 0; i < 2000000; i++ {
		bigi := big.NewInt(int64(i))
		if bigi.ProbablyPrime(4) {
			sum += i
		}
	}
	return sum
}

func main() {
	result := problem()
	fmt.Println("Answer:", result)
}
