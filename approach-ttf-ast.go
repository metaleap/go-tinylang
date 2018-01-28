package main

import (
	"errors"
	"fmt"
)

// TTF approach --- vs. ADT approach

type iTtfRepr interface {
	fmt.Stringer
}

type iTtfSymantics interface {
	lit(num) iTtfRepr
	neg(iTtfRepr) iTtfRepr
	pos(iTtfRepr) iTtfRepr
	add(iTtfRepr, iTtfRepr) iTtfRepr
	sub(iTtfRepr, iTtfRepr) iTtfRepr
	mul(iTtfRepr, iTtfRepr) iTtfRepr
	div(iTtfRepr, iTtfRepr) (iTtfRepr, error)
}

var (
	testnum iTtfSymantics = ttfSymNum{}
	teststr iTtfSymantics = ttfSymStr{}
)

type ttfReprNum num

func (me ttfReprNum) String() string { return num(me).String() }

type ttfSymNum struct{}

func (me ttfSymNum) lit(n num) iTtfRepr {
	return ttfReprNum(n)
}

func (me ttfSymNum) neg(r iTtfRepr) iTtfRepr {
	return -(r.(ttfReprNum))
}

func (me ttfSymNum) pos(r iTtfRepr) iTtfRepr {
	return +(r.(ttfReprNum))
}

func (me ttfSymNum) add(l iTtfRepr, r iTtfRepr) iTtfRepr {
	return l.(ttfReprNum) + r.(ttfReprNum)
}

func (me ttfSymNum) sub(l iTtfRepr, r iTtfRepr) iTtfRepr {
	return l.(ttfReprNum) - r.(ttfReprNum)
}

func (me ttfSymNum) mul(l iTtfRepr, r iTtfRepr) iTtfRepr {
	return l.(ttfReprNum) + r.(ttfReprNum)
}

func (me ttfSymNum) div(l iTtfRepr, r iTtfRepr) (iTtfRepr, error) {
	left, right := l.(ttfReprNum), r.(ttfReprNum)
	if right == 0 {
		return right, errors.New("division of " + l.String() + " by zero")
	}
	return left / right, nil
}

type ttfReprStr string

func (me ttfReprStr) String() string { return string(me) }

type ttfSymStr struct{}

func (me ttfSymStr) lit(n num) iTtfRepr {
	return ttfReprStr(n.String())
}

func (me ttfSymStr) neg(r iTtfRepr) iTtfRepr {
	return ttfReprStr("(-" + r.String() + ")")
}

func (me ttfSymStr) pos(r iTtfRepr) iTtfRepr {
	return ttfReprStr("(+" + r.String() + ")")
}

func (me ttfSymStr) add(l iTtfRepr, r iTtfRepr) iTtfRepr {
	return ttfReprStr("(" + l.String() + " + " + r.String() + ")")
}

func (me ttfSymStr) sub(l iTtfRepr, r iTtfRepr) iTtfRepr {
	return ttfReprStr("(" + l.String() + " - " + r.String() + ")")
}

func (me ttfSymStr) mul(l iTtfRepr, r iTtfRepr) iTtfRepr {
	return ttfReprStr("(" + l.String() + " * " + r.String() + ")")
}

func (me ttfSymStr) div(l iTtfRepr, r iTtfRepr) (iTtfRepr, error) {
	return ttfReprStr("(" + l.String() + " / " + r.String() + ")"), nil
}
