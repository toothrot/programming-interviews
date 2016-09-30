package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	queue := []string{}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines[1:] {
		commands := strings.Split(line, " ")
		oper := commands[0]
		switch oper {
		case "1":
			queue = append(queue, commands[1])
		case "2":
			queue = queue[1:]
		case "3":
			fmt.Println(queue[0])
		}
	}
}
