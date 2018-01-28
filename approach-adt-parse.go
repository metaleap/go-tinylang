package main

import (
	"errors"
)

func adtParse(src string) (expr iAdtExpr, err error) {
	var tokens []token
	if tokens, err = lex(src); err == nil {
		var last iAdtExpr
		var stack []iAdtExpr
		for _, token := range tokens {
			var cur iAdtExpr
			switch xtok := token.(type) {
			case tokenNum:
				cur = adtNum(float64(xtok))
			case tokenOpMinus:
				cur = adtOp2(nil, "-", nil)
			case tokenOpPlus:
				cur = adtOp2(nil, "+", nil)
			case tokenOpSlash:
				cur = adtOp2(nil, "/", nil)
			case tokenOpTimes:
				cur = adtOp2(nil, "*", nil)
			case tokenSepParenOpen:
				stack = append(stack, last)
				last = nil
			case tokenSepParenClose:
				if len(stack) == 0 {
					err = errors.New("mis-matched opening and closing parens")
				} else {
					if last == nil {
						err = errors.New("empty parens")
					} else {
						outer := stack[len(stack)-1]
						if outer1, ok1 := outer.(*adtExprOp1); ok1 {
							outer1.Right = last
							last, cur = nil, outer1
						} else if outer2, ok2 := outer.(*adtExprOp2); ok2 {
							outer2.Right = last
							last, cur = nil, outer2
						} else if outer != nil {
							err = errors.New("invalid left-hand-side of parens: " + outer.String())
						}
					}
					stack = stack[:len(stack)-1]
				}
			}
			if cur != nil {
				if last == nil {
					last = cur
				} else if last, err = cur.parseJoinPrev(last); err != nil {
					last = nil
					break
				} else if last == nil {
					err = errors.New("parse error")
					break
				}
			}
			if err != nil {
				break
			}
		}
		if len(stack) > 0 {
			err = errors.New("mis-matched opening and closing parens")
		}
		if err == nil {
			expr = last
		}
	}
	return
}

func (me *adtExprNum) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	switch xp := prev.(type) {
	case *adtExprOp2:
		if xp.Left == nil {
			expr = adtOp1(xp.Op, me)
		} else {
			xp.Right = me
			expr = xp
		}
	}
	if expr == nil {
		err = errors.New("invalid symbol preceding " + me.String())
	}
	return
}

func (me *adtExprOp1) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	err = errors.New("unexpected left-hand side of unary operator: " + me.Op)
	return
}

func (me *adtExprOp2) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	me.Left = prev
	expr = me
	return
}
