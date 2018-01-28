package main

import (
	"fmt"
)

type num float64 // just for fmt.Stringer.String()

func (me num) String() string { return fmt.Sprintf("%g", float64(me)) }

type iAdtExpr interface {
	fmt.Stringer
	parseJoinPrev(iAdtExpr) (iAdtExpr, error)
}

type adtExprNum struct {
	Num num
}

type adtExprOp1 struct {
	Expr iAdtExpr
}

type adtExprOp2 struct {
	Left  iAdtExpr
	Right iAdtExpr
}

type adtExprOp1Neg struct{ adtExprOp1 }
type adtExprOp1Pos struct{ adtExprOp1 }
type adtExprOp2Add struct{ adtExprOp2 }
type adtExprOp2Sub struct{ adtExprOp2 }
type adtExprOp2Mul struct{ adtExprOp2 }
type adtExprOp2Div struct{ adtExprOp2 }

func adtNum(n float64) iAdtExpr {
	return adtExprNum{Num: num(n)}
}

func (me adtExprNum) String() string { return me.Num.String() }

func adtOpNeg(expr iAdtExpr) iAdtExpr {
	return adtExprOp1Neg{adtExprOp1: adtExprOp1{Expr: expr}}
}

func (me adtExprOp1Neg) String() string { return fmt.Sprintf("(-%s)", me.Expr) }

func adtOpPos(expr iAdtExpr) iAdtExpr {
	return adtExprOp1Pos{adtExprOp1: adtExprOp1{Expr: expr}}
}

func (me adtExprOp1Pos) String() string { return fmt.Sprintf("(+%s)", me.Expr) }

func adtOpAdd(left iAdtExpr, right iAdtExpr) iAdtExpr {
	return adtExprOp2Add{adtExprOp2: adtExprOp2{Left: left, Right: right}}
}

func (me adtExprOp2Add) String() string { return fmt.Sprintf("(%s + %s)", me.Left, me.Right) }

func adtOpSub(left iAdtExpr, right iAdtExpr) iAdtExpr {
	return adtExprOp2Sub{adtExprOp2: adtExprOp2{Left: left, Right: right}}
}

func (me adtExprOp2Sub) String() string { return fmt.Sprintf("(%s - %s)", me.Left, me.Right) }

func adtOpMul(left iAdtExpr, right iAdtExpr) iAdtExpr {
	return adtExprOp2Mul{adtExprOp2: adtExprOp2{Left: left, Right: right}}
}

func (me adtExprOp2Mul) String() string { return fmt.Sprintf("(%s * %s)", me.Left, me.Right) }

func adtOpDiv(left iAdtExpr, right iAdtExpr) iAdtExpr {
	return adtExprOp2Div{adtExprOp2: adtExprOp2{Left: left, Right: right}}
}

func (me adtExprOp2Div) String() string { return fmt.Sprintf("(%s / %s)", me.Left, me.Right) }
