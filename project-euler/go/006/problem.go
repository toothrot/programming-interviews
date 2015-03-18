package main

import "fmt"

func sumOfSquares() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += (i * i)
	}
	return sum
}

func squareOfSums() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return (sum * sum)
}

func problem() int {
	return squareOfSums() - sumOfSquares()
}

func main() {
	result := problem()
	fmt.Println("Answer:", result)
}
