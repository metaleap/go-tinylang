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
				cur = adtOpSub(nil, nil)
			case tokenOpPlus:
				cur = adtOpAdd(nil, nil)
			case tokenOpSlash:
				cur = adtOpDiv(nil, nil)
			case tokenOpTimes:
				cur = adtOpMul(nil, nil)
			case tokenSepParenOpen:
				stack = append(stack, last)
				last = nil
			case tokenSepParenClose:
				if len(stack) == 0 {
					err = errors.New("mis-matched open/close parens")
				} else if last == nil {
					err = errors.New("empty parens")
				} else {
					outer := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					cur = last
					last = outer
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
			err = errors.New("not all parens were fully closed")
		}
		if err == nil {
			expr = last
		}
	}
	return
}

func (me adtExprOp1Neg) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	err = errors.New("the unexpected occurred --- bug in parser!")
	return
}

func (me adtExprOp1Pos) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	err = errors.New("the unexpected occurred --- bug in parser!")
	return
}

func (me adtExprNum) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	switch xp := prev.(type) {
	case adtExprOp2Add:
		if xp.Left == nil {
			expr = adtOpPos(me)
		} else {
			xp.Right = me
			expr = xp
		}
	case adtExprOp2Sub:
		if xp.Left == nil {
			expr = adtOpNeg(me)
		} else {
			xp.Right = me
			expr = xp
		}
	case adtExprOp2Mul:
		xp.Right = me
		expr = xp
	case adtExprOp2Div:
		xp.Right = me
		expr = xp
	}
	if expr == nil {
		err = errors.New("invalid symbol preceding " + me.String())
	}
	return
}

func (me adtExprOp2Add) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	me.Left = prev
	expr = me
	return
}

func (me adtExprOp2Sub) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	me.Left = prev
	expr = me
	return
}

func (me adtExprOp2Mul) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	me.Left = prev
	expr = me
	return
}

func (me adtExprOp2Div) parseJoinPrev(prev iAdtExpr) (expr iAdtExpr, err error) {
	me.Left = prev
	expr = me
	return
}
