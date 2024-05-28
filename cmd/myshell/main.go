package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func run(out io.Writer, cmd string, args ...string) {
	switch cmd {
	case "exit":
		var (
			code int
			err  error
		)
		if len(args) > 0 {
			code, err = strconv.Atoi(args[0])
			if err != nil {
				panic(fmt.Sprintf("exit: invalid code argument: %s", args[0]))
			}
		}

		os.Exit(code)

	case "echo":
		fmt.Fprintf(out, "%s\n", strings.Join(args, " "))

	default:
		fmt.Fprintf(out, "%s: command not found\n", cmd)
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	for {
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

		run(os.Stdout, split[0], split[1:]...)
	}
}
