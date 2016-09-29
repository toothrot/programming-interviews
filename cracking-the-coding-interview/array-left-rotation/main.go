package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(b), "\n")

	args := strings.Split(lines[0], " ")
	input := strings.Split(lines[1], " ")
	size, _ := strconv.ParseInt(args[0], 10, 64)
	rotations, _ := strconv.ParseInt(args[1], 10, 64)
	rotations = rotations % size

	output := append(input[rotations:], input[:rotations]...)
	fmt.Printf("%s\n", strings.Join(output, " "))
}
