package main

import (
	"errors"
)

func parse(tokenStream []iToken) (expr iExpr, err error) {
	var prev iExpr
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
				stack = append(stack, prev)
				prev = nil
			} else if len(stack) == 0 {
				err = errors.New("excess closing paren")
			} else {
				if prev == nil {
					err = errors.New("empty parens")
				} else {
					outer := stack[len(stack)-1]
					if outer1, ok1 := outer.(*exprOp1); ok1 {
						outer1.Right = prev
						prev, cur = nil, outer1
					} else if outer2, ok2 := outer.(*exprOp2); ok2 {
						if outer2.Right = prev; outer2.Left == nil {
							prev, cur = nil, newOp1(outer2.Op, outer2.Right)
						} else {
							prev, cur = nil, outer2
						}
					} else if outer != nil {
						err = errors.New("invalid left-hand-side of parens: " + outer.String())
					}
				}
				stack = stack[:len(stack)-1]
			}
		} // switch-case token.(type)

		if err == nil && cur != nil {
			if prev == nil {
				prev = cur
			} else if prev, err = cur.joinPreceding(prev); err != nil {
				prev = nil
				break
			} else if prev == nil {
				err = errors.New("bug in parser: joinPreceding returned nil,nil")
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
		expr = prev
	}
	return
}

func (this *exprOp1) joinPreceding(prev iExpr) (expr iExpr, err error) {
	err = errors.New("unexpected left-hand side of unary operator: " + this.Op)
	return
}

func (this *exprOp2) joinPreceding(prev iExpr) (expr iExpr, err error) {
	if op2, _ := prev.(*exprOp2); op2 != nil && op2.Left == nil && op2.Right == nil {
		expr = newOp1(op2.Op, this)
	} else {
		this.Left = prev
		expr = this
	}
	return
}

func (this *exprLit) joinPreceding(prev iExpr) (expr iExpr, err error) {
	switch prevop := prev.(type) {
	case *exprOp1:
		if prevop.Right == nil {
			prevop.Right = this
			expr = prevop
		}
	case *exprOp2:
		if prevop.Left == nil {
			expr = newOp1(prevop.Op, this)
		} else if prevop.Right == nil {
			prevop.Right = this
			expr = prevop
		}
	}
	if expr == nil {
		err = errors.New(this.String() + " cannot follow " + str(prev, "?").String())
	}
	return
}
