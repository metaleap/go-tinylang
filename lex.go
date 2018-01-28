package main

import (
	"errors"
	"strconv"
	"strings"
	"text/scanner"
)

type iToken interface{}
type tokenNum float64
type tokenOp string
type tokenParen bool // true for opening, false for closing paren .. hacky, eh?

func lex(src string) (tokenStream []iToken, err error) { // accumulating in a return slice defeats the idea of scalable streaming (eg. parallel lex-parse pipeline), but really no matter in this toy
	var lexer scanner.Scanner
	lexer.Init(strings.NewReader(src)).Filename = src
	lexer.Mode = scanner.ScanFloats | scanner.ScanInts | scanner.SkipComments
	lexer.Error = func(_ *scanner.Scanner, msg string) { err = errors.New(msg) }
	for tok := lexer.Scan(); (err == nil) && (tok != scanner.EOF); tok = lexer.Scan() {
		if sym := lexer.TokenText(); sym != "" { // should never be "" anyway, really
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
					tokenStream = append(tokenStream, tokenOp(sym))
				case "(", ")":
					tokenStream = append(tokenStream, tokenParen(sym == "("))
				default:
					err = errors.New("unrecognized token: " + sym)
				}
			}
		}
	}
	return
}
