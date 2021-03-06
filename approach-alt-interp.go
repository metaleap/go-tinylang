package main

import (
	"fmt"
)

func altInterp_PrettyPrint(expr iExpr) (fmt.Stringer, error) {
	s, e := altSymStr{}.interp(expr)
	return s.(altReprStr), e
}

func altInterp_Eval(expr iExpr) (fmt.Stringer, error) {
	n, e := altSymNum{}.interp(expr)
	return n.(altReprNum), e
}

// representation type
type iAltRepr interface{}

type altReprNum num

func (this altReprNum) String() string { return num(this).String() }

type altReprStr string

func (this altReprStr) String() string { return string(this) }

// language vocabulary
type iAltSymantics interface {
	interp(iExpr) (iAltRepr, error)
	lit(*exprLit) iAltRepr

	neg(iAltRepr) iAltRepr
	pos(iAltRepr) iAltRepr
	add(iAltRepr, iAltRepr) iAltRepr
	sub(iAltRepr, iAltRepr) iAltRepr
	mul(iAltRepr, iAltRepr) iAltRepr
	div(iAltRepr, iAltRepr) (iAltRepr, error)
}

type altSymNum struct{}

func (altSymNum) neg(r iAltRepr) iAltRepr             { return -(r.(altReprNum)) }
func (altSymNum) pos(r iAltRepr) iAltRepr             { return +(r.(altReprNum)) }
func (altSymNum) add(l iAltRepr, r iAltRepr) iAltRepr { return l.(altReprNum) + r.(altReprNum) }
func (altSymNum) sub(l iAltRepr, r iAltRepr) iAltRepr { return l.(altReprNum) - r.(altReprNum) }
func (altSymNum) mul(l iAltRepr, r iAltRepr) iAltRepr { return l.(altReprNum) * r.(altReprNum) }
func (altSymNum) div(l iAltRepr, r iAltRepr) (iAltRepr, error) {
	left, right := l.(altReprNum), r.(altReprNum)
	if right == 0 {
		return altReprNum(0), errInterpDiv0(left.String())
	}
	return left / right, nil
}

func (this altSymNum) lit(numLit *exprLit) iAltRepr {
	var n num
	if numLit != nil {
		n = numLit.Num
	}
	return altReprNum(n)
}

func (this altSymNum) interp(expr iExpr) (iAltRepr, error) {
	return _altInterp(this, expr)
}

type altSymStr struct{}

// altSymStr ignores that iExpr implements fmt.Stringer (or pretends that it might not; or does not require it to) ——— for demo reasons

func (altSymStr) neg(r iAltRepr) iAltRepr { return "(-" + r.(altReprStr) + ")" }
func (altSymStr) pos(r iAltRepr) iAltRepr { return "(+" + r.(altReprStr) + ")" }
func (altSymStr) add(l iAltRepr, r iAltRepr) iAltRepr {
	return "(" + l.(altReprStr) + " + " + r.(altReprStr) + ")"
}
func (altSymStr) sub(l iAltRepr, r iAltRepr) iAltRepr {
	return "(" + l.(altReprStr) + " - " + r.(altReprStr) + ")"
}
func (altSymStr) mul(l iAltRepr, r iAltRepr) iAltRepr {
	return "(" + l.(altReprStr) + " * " + r.(altReprStr) + ")"
}
func (altSymStr) div(l iAltRepr, r iAltRepr) (iAltRepr, error) {
	return "(" + l.(altReprStr) + " / " + r.(altReprStr) + ")", nil
}

func (altSymStr) lit(numLit *exprLit) iAltRepr {
	var s string
	if numLit != nil {
		s = numLit.Num.String()
	}
	return altReprStr(s)
}

func (this altSymStr) interp(expr iExpr) (iAltRepr, error) {
	return _altInterp(this, expr)
}

func _altInterp(me iAltSymantics, expr iExpr) (repr iAltRepr, err error) {
	switch x := expr.(type) {
	case *exprLit:
		repr = me.lit(x)
	case *exprOp1:
		repr, err = _altInterpOp1(me, x)
	case *exprOp2:
		repr, err = _altInterpOp2(me, x)
	}
	if err != nil {
		repr = me.lit(nil)
	}
	return
}

func _altInterpOp1(me iAltSymantics, expr *exprOp1) (repr iAltRepr, err error) {
	var fn func(iAltRepr) iAltRepr
	switch expr.Op {
	case "+":
		fn = me.pos
	case "-":
		fn = me.neg
	}
	if fn == nil {
		err = errInterpBadOp1(expr.Op)
	} else if expr.Right == nil {
		err = errInterpLate(expr)
	} else if repr, err = me.interp(expr.Right); err == nil {
		repr = fn(repr)
	} else {
		repr = nil
	}
	return
}

func _altInterpOp2(me iAltSymantics, expr *exprOp2) (repr iAltRepr, err error) {
	var fn func(iAltRepr, iAltRepr) iAltRepr
	switch expr.Op {
	case "+":
		fn = me.add
	case "-":
		fn = me.sub
	case "*":
		fn = me.mul
	case "/":
		fn = func(l iAltRepr, r iAltRepr) (rep iAltRepr) {
			rep, err = me.div(l, r)
			return
		}
	}
	var left, right iAltRepr
	if fn == nil {
		err = errInterpBadOp2(expr.Op)
	} else if expr.Left == nil || expr.Right == nil {
		err = errInterpLate(expr)
	} else if left, err = me.interp(expr.Left); err == nil {
		if right, err = me.interp(expr.Right); err == nil {
			repr = fn(left, right)
		}
	}
	return
}
