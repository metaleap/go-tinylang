package main

import (
	"bufio"
	"os"

	"github.com/go-leap/str"
)

var writeln = os.Stdout.WriteString

func main() {
	repl := bufio.NewScanner(os.Stdin)
	writeln(`REPL for the ad-hoc NanoCalc language, consisting only of number operands and some arithmetic operators:

- :q to quit
- <expr> to parse-and-eval

`)
ReadEvalPrintLoop:
	for repl.Scan() {
		if err := repl.Err(); err != nil {
			panic(err)
		} else if readln := ustr.Trim(repl.Text()); readln != "" {
			switch readln {
			case ":q":
				break ReadEvalPrintLoop
			default:
				if err = parseAndEval(readln); err != nil {
					println(err)
				}
			}
		}
	}
}

func parseAndEval(expr string) (err error) {
	return
}
