package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aeharvlee/compiler-dfa/dfa/hard_wired"
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return input
}

func main() {
	for {
		input := readInput()
		hardWiredDFA := hard_wired.NewHardWiredDFA()
		if hardWiredDFA.Recognize(input) {
			fmt.Println("Accepted.")
		} else {
			fmt.Println("Not Accepted.")
		}
	}
}
