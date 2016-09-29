package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(b), "\n")
	line1 := strings.Split(lines[1], " ")
	line2 := strings.Split(lines[2], " ")

	line1map := make(map[string]int)

	for _, word := range line1 {
		line1map[word] += 1
	}

	for _, word := range line2 {
		v, ok := line1map[word]
		if !ok {
			fmt.Println("No")
			return
		}
		v -= 1
		if v < 0 {
			fmt.Println("No")
			return
		}
		line1map[word] = v
	}
	fmt.Println("Yes")
}
