package main

import (
	"fmt"
)

type iExpr interface {
	fmt.Stringer
	joinPreceding(iExpr) (iExpr, error)
}

func newLit(n float64) iExpr                   { return &exprLit{Num: num(n)} }
func newOp1(op string, r iExpr) iExpr          { return &exprOp1{Op: op, Right: r} }
func newOp2(l iExpr, op string, r iExpr) iExpr { return &exprOp2{Left: l, Op: op, Right: r} }

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

func (me *exprLit) String() string { return me.Num.String() }
func (me *exprOp1) String() string { return strf("(%s%s)", me.Op, str(me.Right, "?")) }
func (me *exprOp2) String() string {
	return strf("(%s %s %s)", str(me.Left, "?"), me.Op, str(me.Right, "?"))
}
