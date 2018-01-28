package main

import (
	"fmt"
)

type iExpr interface {
	fmt.Stringer
	parseJoinPrev(iExpr) (iExpr, error)
}

type exprLit struct {
	Num num
}

type exprOp1 struct {
	Op    string
	Right iExpr
}

type exprOp2 struct {
	Left  iExpr
	Op    string
	Right iExpr
}

func newLit(n float64) iExpr {
	return &exprLit{Num: num(n)}
}

func (me *exprLit) String() string { return me.Num.String() }

func newOp1(op string, right iExpr) iExpr {
	return &exprOp1{Op: op, Right: right}
}

func (me *exprOp1) String() string { return fmt.Sprintf("(%s%s)", me.Op, stringer(me.Right)) }

func newOp2(left iExpr, op string, right iExpr) iExpr {
	return &exprOp2{Left: left, Op: op, Right: right}
}

func (me *exprOp2) String() string {
	return fmt.Sprintf("(%s %s %s)", stringer(me.Left), me.Op, stringer(me.Right))
}
