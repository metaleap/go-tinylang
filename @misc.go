package main

import (
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

func stringer(str interface{}) (s fmt.Stringer) {
	s, _ = str.(fmt.Stringer)
	if s == nil {
		s = strNil{}
	}
	return
}

type strNil struct{}

func (strNil) String() string { return "?" }
