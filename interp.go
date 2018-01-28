package main

import (
	"fmt"
)

type adtInterp func(expr iAdtExpr) (fmt.Stringer, error)

func adtInterp_Eval(expr iAdtExpr) (fmt.Stringer, error) {
	switch me := expr.(type) {
	case adtExprNum:
		return num(me.Num), nil
	case adtExprOp1Pos:
		n, e := adtInterp_Eval(me.Expr)
		return +(n.(num)), e
	case adtExprOp1Neg:
		n, e := adtInterp_Eval(me.Expr)
		return -(n.(num)), e
	case adtExprOp2Add:
		n1, e1 := adtInterp_Eval(me.Left)
		n2, e2 := adtInterp_Eval(me.Right)
		return n1.(num) + n2.(num), errPick(e1, e2)
	case adtExprOp2Sub:
		n1, e1 := adtInterp_Eval(me.Left)
		n2, e2 := adtInterp_Eval(me.Right)
		return n1.(num) - n2.(num), errPick(e1, e2)
	case adtExprOp2Mul:
		n1, e1 := adtInterp_Eval(me.Left)
		n2, e2 := adtInterp_Eval(me.Right)
		return n1.(num) * n2.(num), errPick(e1, e2)
	case adtExprOp2Div:
		n1, e1 := adtInterp_Eval(me.Left)
		n2, e2 := adtInterp_Eval(me.Right)
		if n2.(num) == 0.0 {
			return num(0), fmt.Errorf("Division by zero in: %s", expr)
		}
		return n1.(num) / n2.(num), errPick(e1, e2)
	}
	return num(0), fmt.Errorf("Invalid form: %s", expr)
}

func adtInterp_PrettyPrint(expr iAdtExpr) (fmt.Stringer, error) {
	return expr, nil
}
