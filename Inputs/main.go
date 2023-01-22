package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello Developer")
	reader := bufio.NewReader(os.Stdin)
	var stop bool = false
	for !stop {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "Exit" {
			stop = true
		}
		println(stop)

	}
}
