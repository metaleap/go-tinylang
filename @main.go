package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const taglessFinal = false

type num float64 // just for fmt.Stringer.String()

func (me num) String() string { return fmt.Sprintf("%g", float64(me)) }

func writeLn(s string) { _, _ = os.Stdout.WriteString(s + "\n") }

func main() {
	repl := bufio.NewScanner(os.Stdin)
	writeLn(`REPL for our mini 'NanoCalc'
language, consisting only of:
float operands, parens and the
most basic arithmetic operators
(of equal precedence: use parens).

Type:
· Q to quit
· <expr> to parse-and-prettyprint-and-eval

`)
	for repl.Scan() {
		if err := repl.Err(); err != nil {
			panic(err)
		} else if readln := strings.TrimSpace(repl.Text()); readln != "" {
			switch readln {
			case "q", "Q":
				return
			default:
				if taglessFinal {
					err = errors.New("TODO: ttf approach")
				} else {
					err = adtParseAndInterp(readln, adtInterp_PrettyPrint, adtInterp_Eval)
				}
				if err != nil {
					println(err.Error())
				}
			}
		}
	}
}

func adtParseAndInterp(src string, interps ...adtInterp) (err error) {
	var expr iAdtExpr
	var val fmt.Stringer
	if expr, err = adtParse(src); err == nil {
		for _, interp := range interps {
			if val, err = interp(expr); err != nil {
				break
			} else {
				fmt.Printf("\n%s\n", val)
			}
		}
	}
	return
}

func errPick(errs ...error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

func stringer(str interface{}) (s fmt.Stringer) {
	s, _ = str.(fmt.Stringer)
	if s == nil {
		s = strNil{}
	}
	return
}

type strNil struct{}

func (strNil) String() string { return "?" }
