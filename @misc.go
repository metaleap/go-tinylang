package main

import (
	"errors"
	"fmt"
	"os"
)

type interpreter func(expr iExpr) (fmt.Stringer, error)

type num float64

func (me num) String() string { return strf("%g", float64(me)) }

var strf = fmt.Sprintf

func writeLn(s string) { _, _ = os.Stdout.WriteString(s + "\n") }

func errPick(errs ...error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

func errInterpBadOp1(op string) error {
	return errors.New("invalid unary operator: " + op)
}

func errInterpBadOp2(op string) error {
	return errors.New("invalid binary operator: " + op)
}

func errInterpDiv0(n string) error {
	return errors.New("division of " + n + " by zero")
}

func errInterpLate(expr iExpr) error {
	return errors.New("invalid operand or operator in: " + expr.String())
}

func str(any interface{}) (stringer fmt.Stringer) {
	stringer, _ = any.(fmt.Stringer)
	if stringer == nil {
		stringer = strNil{}
	}
	return
}

type strNil struct{}

func (strNil) String() string { return "?" }
