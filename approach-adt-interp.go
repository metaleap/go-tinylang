package main

import (
	"fmt"
)

func adtInterp_PrettyPrint(expr iExpr) (fmt.Stringer, error) {
	return expr, nil
}

func adtInterp_Eval(expr iExpr) (fmt.Stringer, error) {
	switch me := expr.(type) {
	case *exprLit:
		return num(me.Num), nil
	case *exprOp1:
		if me.Right != nil {
			n, e := adtInterp_Eval(me.Right)
			switch me.Op {
			case "+":
				return +(n.(num)), e
			case "-":
				return -(n.(num)), e
			default:
				return num(0), errInterpBadOp1(me.Op)
			}
		}
	case *exprOp2:
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
					return num(0), errPick(e1, e2, errInterpDiv0(n1.String()))
				}
				return n1.(num) / n2.(num), errPick(e1, e2)
			default:
				return num(0), errInterpBadOp2(me.Op)
			}
		}
	}
	return num(0), errInterpLate(expr)
}
