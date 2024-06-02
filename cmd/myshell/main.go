package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

var pathEntries []string

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
}

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
		fmt.Fprintln(out, strings.Join(args, " "))

	case "type":
		for _, a := range args {
			if builtins[a] {
				fmt.Fprintf(out, "%s is a shell builtin\n", a)
				continue
			}

			fullPath, ok := lookupBinary(a)
			if ok {
				fmt.Fprintf(out, "%s is %s\n", a, fullPath)
				continue
			}

			fmt.Fprintf(out, "%s not found\n", a)
		}

	case "pwd":
		wd, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("can't get current working directory: %v", err))
		}

		fmt.Fprintln(out, wd)

	case "":
		return

	default:
		command := exec.Command(cmd, args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			fmt.Fprintf(out, "%s: command not found\n", command)
		}
	}
}

func lookupBinary(name string) (fullPath string, ok bool) {
	for _, p := range pathEntries {
		fp := path.Join(p, name)

		_, err := os.Stat(fp)
		if err == nil {
			return fp, true
		}
	}

	return
}

func main() {
	path := os.Getenv("PATH")
	pathEntries = strings.Split(path, ":")

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
