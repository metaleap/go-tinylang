package main

import (
	"errors"
	"strconv"
	"strings"
	"text/scanner"
)

type token interface{}

type tokenNum float64

type tokenOpPlus struct{}

type tokenOpMinus struct{}

type tokenOpTimes struct{}

type tokenSepParenClose struct{}

type tokenSepParenOpen struct{}

func lex(src string) (tokens []token, err error) { // accumulating in a return slice defeats the idea of scalable streaming (eg. parallel lex-parse pipeline), but really no matter in this toy
	var lex scanner.Scanner
	lex.Init(strings.NewReader(src)).Filename = src
	lex.Mode = scanner.ScanFloats | scanner.ScanInts | scanner.SkipComments
	lex.Error = func(_ *scanner.Scanner, msg string) { err = errors.New(msg) }
	for tok := lex.Scan(); (err == nil) && (tok != scanner.EOF); tok = lex.Scan() {
		if sym := lex.TokenText(); sym != "" { // should never be "", but we're defensive for once
			switch tok {
			case scanner.Float:
				var f float64
				if f, err = strconv.ParseFloat(sym, 64); err == nil {
					tokens = append(tokens, tokenNum(f))
				}
			case scanner.Int:
				var i int64
				if i, err = strconv.ParseInt(sym, 0, 64); err == nil {
					tokens = append(tokens, tokenNum(i))
				}
			default:
				switch sym {
				case "+":
					tokens = append(tokens, tokenOpPlus{})
				case "-":
					tokens = append(tokens, tokenOpMinus{})
				case "*":
					tokens = append(tokens, tokenOpTimes{})
				case "(":
					tokens = append(tokens, tokenSepParenOpen{})
				case ")":
					tokens = append(tokens, tokenSepParenClose{})
				default:
					err = errors.New("Unrecognized symbol: " + sym)
				}
			}
		}
	}
	return
}
