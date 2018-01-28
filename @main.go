package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	interp_prettyprint, interp_eval := adtInterp_PrettyPrint, adtInterp_Eval
	alt, repl := false, bufio.NewScanner(os.Stdin)

	writeLn(`
——————————————————————————————————————————————
  REPL for our demo 'TinyCalc'
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
——————————————————————————————————————————————
`)

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
					writeLn("Now using 'Alt' interpreter approach")
					interp_prettyprint, interp_eval = altInterp_PrettyPrint, altInterp_Eval
				} else {
					writeLn("Now using 'ADT' interpreter approach")
					interp_prettyprint, interp_eval = adtInterp_PrettyPrint, adtInterp_Eval
				}
			default:
				if err = lexAndParseAndInterp(readln, interp_prettyprint, interp_eval); err != nil {
					println(err.Error())
				}
			}
		}
	}
}

func lexAndParseAndInterp(src string, interps ...interpreter) (err error) {
	var lexed []iToken
	if lexed, err = lex(src); err == nil {
		var expr iExpr
		if expr, err = parse(lexed); err == nil {
			for _, interp := range interps {
				var val fmt.Stringer
				if val, err = interp(expr); err != nil {
					break
				} else {
					writeLn(val.String())
				}
			}
		}
	}
	return
}
