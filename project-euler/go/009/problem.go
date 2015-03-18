package main

import "fmt"
import "math"

func problem() int {
	for a := 0; a < 1000; a++ {
		for b := a + 1; b < 1000-a; b++ {
			cSquare := (a * a) + (b * b)
			c := int(math.Sqrt(float64(cSquare)))
			perfectSquare := math.Sqrt(float64(cSquare)) == float64(int(math.Sqrt(float64(cSquare))))

			if perfectSquare && c > b && (a+b+c) == 1000 {
				fmt.Println("Answer:", a, b, c)
				return a * b * c
			}
		}
	}
	return 0
}

func main() {
	result := problem()
	fmt.Println("Answer:", result)
}
