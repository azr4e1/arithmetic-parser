package main

import "errors"

type Lexer interface {
	Advance()
	NextToken() Terminal
}

type ArithmeticLexer struct {
	Pos     int
	Current rune
	Text    string
}

func NewAL(text string) *ArithmeticLexer {
	if len(text) == 0 {
		return nil
	}

	return &ArithmeticLexer{
		Text:    text,
		Current: rune(text[0]),
	}
}

func (al *ArithmeticLexer) Advance() {
	al.Pos++
	if al.Pos < len(al.Text) {
		al.Current = rune(al.Text[al.Pos])
	} else {
		al.Current = *new(rune)
	}
}

func (al *ArithmeticLexer) SkipSpace() {
	for al.Current == ' ' || al.Current == '\t' {
		al.Current = rune(al.Text[al.Pos])
		al.Advance()
	}
}

func (al *ArithmeticLexer) NextToken() (Terminal, error) {
	if term, ok := Digits[al.Current]; ok {
		al.Advance()
		return term, nil
	}
	if term, ok := Symbols[al.Current]; ok {
		if
		return Zero, errors.New("Symbol not recognized by vocabulary.")
	}
	return Zero, nil
}
