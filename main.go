package main

import (
	"bufio"
	"os"

	"github.com/go-leap/str"
)

func main() {
	writeln, stdin := os.Stdout.WriteString, bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if readln := ustr.Trim(stdin.Text()); readln != "" {
			writeln(readln)
		}
	}
}
