package main

import (
	"bufio"
	"os"
	"strings"
)

func writeLn(s string) {
	_, _ = os.Stdout.WriteString(s + "\n")
}

func main() {
	repl := bufio.NewScanner(os.Stdin)
	writeLn(`REPL for the ad-hoc NanoCalc language, consisting only of float operands, parens and the most basic arithmetic operators:

- :q to quit
- <expr> to parse-and-eval

`)
ReadEvalPrintLoop:
	for repl.Scan() {
		if err := repl.Err(); err != nil {
			panic(err)
		} else if readln := strings.TrimSpace(repl.Text()); readln != "" {
			switch readln {
			case ":q":
				break ReadEvalPrintLoop
			default:
				if err = parseAndEval(readln); err != nil {
					println(err.Error())
				}
			}
		}
	}
}

func parseAndEval(src string) (err error) {
	var expr *expr
	if expr, err = parse(src); err == nil {
		if expr, err = eval(expr); err == nil {
			println(expr)
		}
	}
	return
}
