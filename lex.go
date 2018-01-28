package main

import (
	"errors"
	"strconv"
	"strings"
	"text/scanner"
)

type iToken interface{}
type tokenNum float64
type tokenOp struct{ op string }
type tokenParen struct{ opening bool }

func lex(src string) (tokenStream []iToken, err error) { // accumulating in a return slice defeats the idea of scalable streaming (eg. parallel lex-parse pipeline), but really no matter in this toy
	var lex scanner.Scanner
	lex.Init(strings.NewReader(src)).Filename = src
	lex.Mode = scanner.ScanFloats | scanner.ScanInts | scanner.SkipComments
	lex.Error = func(_ *scanner.Scanner, msg string) { err = errors.New(msg) }
	for tok := lex.Scan(); (err == nil) && (tok != scanner.EOF); tok = lex.Scan() {
		if sym := lex.TokenText(); sym != "" { // should never be "" anyway, really
			switch tok {
			case scanner.Float:
				var f float64
				if f, err = strconv.ParseFloat(sym, 64); err == nil {
					tokenStream = append(tokenStream, tokenNum(f))
				}
			case scanner.Int:
				var i int64
				if i, err = strconv.ParseInt(sym, 0, 64); err == nil {
					tokenStream = append(tokenStream, tokenNum(i))
				}
			default:
				switch sym {
				case "+", "-", "*", "/":
					tokenStream = append(tokenStream, tokenOp{op: sym})
				case "(", ")":
					tokenStream = append(tokenStream, tokenParen{opening: sym == "("})
				default:
					err = errors.New("Unrecognized symbol: " + sym)
				}
			}
		}
	}
	return
}
