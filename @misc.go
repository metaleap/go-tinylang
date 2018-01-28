package main

import (
	"errors"
	"fmt"
)

type interp func(expr iExpr) (fmt.Stringer, error)

type num float64 // just for fmt.Stringer.String()

func (me num) String() string { return fmt.Sprintf("%g", float64(me)) }

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

func stringer(str interface{}) (s fmt.Stringer) {
	s, _ = str.(fmt.Stringer)
	if s == nil {
		s = strNil{}
	}
	return
}

type strNil struct{}

func (strNil) String() string { return "?" }
