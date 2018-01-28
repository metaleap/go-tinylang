package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeLn(s string) { _, _ = os.Stdout.WriteString(s + "\n") }

func main() {
	repl := bufio.NewScanner(os.Stdin)
	writeLn(`REPL for our demo 'TinyCalc'
language, consisting only of:
float operands, parens and the 4
most basic arithmetic operators
(with no precedence: use parens).

Enter:
· Q to quit
· A to toggle between:
  · "ADT" interpreter approach (default)
  · "Alt" interpreter approach
· <expr> to parse-and-prettyprint-and-eval

`)
	alt, interp_prettyprint, interp_eval := false, adtInterp_PrettyPrint, adtInterp_Eval
	for repl.Scan() {
		if err := repl.Err(); err != nil {
			panic(err)
		} else if readln := strings.TrimSpace(repl.Text()); readln != "" {
			switch readln {
			case "q", "Q":
				writeLn("Bye!")
				return
			case "a", "A":
				if alt = !alt; alt {
					interp_prettyprint, interp_eval = altInterp_PrettyPrint, altInterp_Eval
					writeLn("Now using 'Alt' interpreter approach")
				} else {
					interp_prettyprint, interp_eval = adtInterp_PrettyPrint, adtInterp_Eval
					writeLn("Now using 'ADT' interpreter approach")
				}
			default:
				if err = lexAndParseAndInterp(readln, interp_prettyprint, interp_eval); err != nil {
					println(err.Error())
				}
			}
		}
	}
}

func lexAndParseAndInterp(src string, interps ...interp) (err error) {
	var lexed []iToken
	if lexed, err = lex(src); err == nil {
		var expr iExpr
		if expr, err = parse(lexed); err == nil {
			for _, interp := range interps {
				var val fmt.Stringer
				if val, err = interp(expr); err != nil {
					break
				} else {
					fmt.Printf("\n%s\n", val)
				}
			}
		}
	}
	return
}
