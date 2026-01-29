package main

type Parser interface {
	Parse() any
}

type ArithmeticParser struct {
	lexer Lexer
}

func (ap *ArithmeticParser) Number()
