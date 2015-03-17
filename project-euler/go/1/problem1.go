package main

import "fmt"

func problem1() int {
	sum := 0

	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	result := problem1()
	fmt.Println("Answer:", result)
}
