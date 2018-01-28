package main

import (
	"errors"
	"fmt"
)

func altInterp_Eval(expr iExpr) (fmt.Stringer, error) {
	return altInterp(expr, altSymNum{})
}

func altInterp_PrettyPrint(expr iExpr) (fmt.Stringer, error) {
	return altInterp(expr, altSymStr{})
}

func altInterp(expr iExpr, sym iAltSymantics) (iAltRepr, error) {
	return sym.interp(expr)
}

type iAltRepr interface {
	fmt.Stringer
	repr()
}

type iAltSymantics interface {
	interp(iExpr) (iAltRepr, error)
	lit(num) iAltRepr
	neg(iAltRepr) iAltRepr
	pos(iAltRepr) iAltRepr
	add(iAltRepr, iAltRepr) iAltRepr
	sub(iAltRepr, iAltRepr) iAltRepr
	mul(iAltRepr, iAltRepr) iAltRepr
	div(iAltRepr, iAltRepr) (iAltRepr, error)
}

type altReprNum num

func (altReprNum) repr()             {}
func (me altReprNum) String() string { return num(me).String() }

type altSymNum struct{}

func (me altSymNum) lit(n num) iAltRepr {
	return altReprNum(n)
}

func (me altSymNum) neg(r iAltRepr) iAltRepr {
	return -(r.(altReprNum))
}

func (me altSymNum) pos(r iAltRepr) iAltRepr {
	return +(r.(altReprNum))
}

func (me altSymNum) add(l iAltRepr, r iAltRepr) iAltRepr {
	return l.(altReprNum) + r.(altReprNum)
}

func (me altSymNum) sub(l iAltRepr, r iAltRepr) iAltRepr {
	return l.(altReprNum) - r.(altReprNum)
}

func (me altSymNum) mul(l iAltRepr, r iAltRepr) iAltRepr {
	return l.(altReprNum) * r.(altReprNum)
}

func (me altSymNum) div(l iAltRepr, r iAltRepr) (iAltRepr, error) {
	left, right := l.(altReprNum), r.(altReprNum)
	if right == 0 {
		return right, errors.New("division of " + l.String() + " by zero")
	}
	return left / right, nil
}

func (me altSymNum) interp(expr iExpr) (repr iAltRepr, err error) {
	switch x := expr.(type) {
	case *exprLit:
		repr = altReprNum(x.Num)
	case *exprOp1:
		repr, err = altInterpOp1(me, x)
	case *exprOp2:
		repr, err = altInterpOp2(me, x)
	}
	return
}

type altReprStr string

func (altReprStr) repr()             {}
func (me altReprStr) String() string { return string(me) }

type altSymStr struct{}

func (me altSymStr) lit(n num) iAltRepr {
	return altReprStr(n.String())
}

func (me altSymStr) neg(r iAltRepr) iAltRepr {
	return altReprStr("(-" + r.(altReprStr) + ")")
}

func (me altSymStr) pos(r iAltRepr) iAltRepr {
	return altReprStr("(+" + r.(altReprStr) + ")")
}

func (me altSymStr) add(l iAltRepr, r iAltRepr) iAltRepr {
	return altReprStr("(" + l.(altReprStr) + " + " + r.(altReprStr) + ")")
}

func (me altSymStr) sub(l iAltRepr, r iAltRepr) iAltRepr {
	return altReprStr("(" + l.(altReprStr) + " - " + r.(altReprStr) + ")")
}

func (me altSymStr) mul(l iAltRepr, r iAltRepr) iAltRepr {
	return altReprStr("(" + l.(altReprStr) + " * " + r.(altReprStr) + ")")
}

func (me altSymStr) div(l iAltRepr, r iAltRepr) (iAltRepr, error) {
	return altReprStr("(" + l.(altReprStr) + " / " + r.(altReprStr) + ")"), nil
}

func (me altSymStr) interp(expr iExpr) (repr iAltRepr, err error) {
	switch x := expr.(type) {
	case *exprLit:
		repr = altReprStr(x.String())
	case *exprOp1:
		repr, err = altInterpOp1(me, x)
	case *exprOp2:
		repr, err = altInterpOp2(me, x)
	}
	return
}

func altInterpOp1(me iAltSymantics, expr *exprOp1) (repr iAltRepr, err error) {
	var fn func(iAltRepr) iAltRepr
	switch expr.Op {
	case "+":
		fn = me.pos
	case "-":
		fn = me.neg
	}
	if fn == nil {
		err = errors.New("invalid unary operator: " + expr.Op)
	} else if repr, err = me.interp(expr.Right); err == nil {
		repr = me.pos(repr)
	} else {
		repr = nil
	}
	return
}

func altInterpOp2(me iAltSymantics, expr *exprOp2) (repr iAltRepr, err error) {
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
		err = errors.New("invalid binary operator: " + expr.Op)
	} else if left, err = me.interp(expr.Left); err == nil {
		if right, err = me.interp(expr.Right); err == nil {
			repr = fn(left, right)
		}
	}
	return
}