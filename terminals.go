package main

// type Terminal int

type Digit int
type Symbol int

const (
	Zero Digit = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

const (
	Dot Symbol = iota
	LeftPar
	RightPar
	Plus
	Minus
	Multiply
	Divide
	Space
	Tab
)

var Digits = map[rune]Digit{
	'0': Zero,
	'1': One,
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
}

var Symbols = map[rune]Symbol{
	'.':  Dot,
	'(':  LeftPar,
	')':  RightPar,
	'+':  Plus,
	'-':  Minus,
	'*':  Multiply,
	'/':  Divide,
	' ':  Space,
	'\t': Tab,
}

// type Number
