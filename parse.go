package main

func parse(src string) (expr *expr, err error) {
	var tokens []token
	if tokens, err = lex(src); err == nil {
		for _, token := range tokens {
			switch xtok := token.(type) {
			case tokenNum:
				println(xtok)
			}
		}
	}
	return
}
