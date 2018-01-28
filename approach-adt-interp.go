package main

import (
	"fmt"
)

type adtInterp func(expr iAdtExpr) (fmt.Stringer, error)

func adtInterp_Eval(expr iAdtExpr) (fmt.Stringer, error) {
	switch me := expr.(type) {
	case *adtExprLit:
		return num(me.Num), nil
	case *adtExprOp1:
		if me.Right != nil {
			n, e := adtInterp_Eval(me.Right)
			switch me.Op {
			case "+":
				return +(n.(num)), e
			case "-":
				return -(n.(num)), e
			default:
				return n.(num), fmt.Errorf("invalid unary operator: " + me.Op)
			}
		}
	case *adtExprOp2:
		if me.Left != nil && me.Right != nil {
			n1, e1 := adtInterp_Eval(me.Left)
			n2, e2 := adtInterp_Eval(me.Right)
			switch me.Op {
			case "+":
				return n1.(num) + n2.(num), errPick(e1, e2)
			case "-":
				return n1.(num) - n2.(num), errPick(e1, e2)
			case "*":
				return n1.(num) * n2.(num), errPick(e1, e2)
			case "/":
				if n2.(num) == 0.0 {
					return num(0), errPick(e1, e2, fmt.Errorf("division by zero in: %s", expr))
				}
				return n1.(num) / n2.(num), errPick(e1, e2)
			}
		}
	}
	return num(0), fmt.Errorf("invalid operand or operator in: %s", expr)
}

func adtInterp_PrettyPrint(expr iAdtExpr) (fmt.Stringer, error) {
	return expr, nil
}
