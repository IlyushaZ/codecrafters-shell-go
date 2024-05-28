package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	in, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("can't read input: %v", err))
	}

	in = strings.TrimSpace(in)

	split := strings.Split(in, " ")
	if len(split) < 1 {
		panic(fmt.Sprintf("invalid input len: %d", len(split)))
	}

	fmt.Fprintf(os.Stdout, "%s: command not found\n", split[0])
}
