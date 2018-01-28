package main

import (
	"errors"
)

func parse(tokenStream []iToken) (expr iExpr, err error) {
	var last iExpr
	var stack []iExpr
	for _, token := range tokenStream {
		var cur iExpr
		switch curtoken := token.(type) {
		case tokenNum:
			cur = newLit(float64(curtoken))
		case tokenOp:
			cur = newOp2(nil, string(curtoken), nil)
		case tokenParen:
			if curtoken { // opening paren if true, else closing paren
				stack = append(stack, last)
				last = nil
			} else if len(stack) == 0 {
				err = errors.New("excess closing paren")
			} else {
				if last == nil {
					err = errors.New("empty parens")
				} else {
					outer := stack[len(stack)-1]
					if outer1, ok1 := outer.(*exprOp1); ok1 {
						outer1.Right = last
						last, cur = nil, outer1
					} else if outer2, ok2 := outer.(*exprOp2); ok2 {
						if outer2.Right = last; outer2.Left == nil {
							last, cur = nil, newOp1(outer2.Op, outer2.Right)
						} else {
							last, cur = nil, outer2
						}
					} else if outer != nil {
						err = errors.New("invalid left-hand-side of parens: " + outer.String())
					}
				}
				stack = stack[:len(stack)-1]
			}
		} // switch-case token.(type)

		if err == nil && cur != nil {
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
	if err == nil && len(stack) > 0 {
		err = errors.New("opening/closing parens mis-match")
	}
	if err == nil {
		expr = last
	}
	return
}

func (me *exprLit) parseJoinPrev(prev iExpr) (expr iExpr, err error) {
	switch xp := prev.(type) {
	case *exprOp1:
		if xp.Right == nil {
			xp.Right = me
			expr = xp
		}
	case *exprOp2:
		if xp.Left == nil {
			expr = newOp1(xp.Op, me)
		} else if xp.Right == nil {
			xp.Right = me
			expr = xp
		}
	}
	if expr == nil {
		err = errors.New(me.String() + " cannot follow " + str(prev).String())
	}
	return
}

func (me *exprOp1) parseJoinPrev(prev iExpr) (expr iExpr, err error) {
	err = errors.New("unexpected left-hand side of unary operator: " + me.Op)
	return
}

func (me *exprOp2) parseJoinPrev(prev iExpr) (expr iExpr, err error) {
	if op2, _ := prev.(*exprOp2); op2 != nil && op2.Left == nil && op2.Right == nil {
		expr = newOp1(op2.Op, me)
	} else {
		me.Left = prev
		expr = me
	}
	return
}
