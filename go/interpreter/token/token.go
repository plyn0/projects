package token

import (
	"errors"
	"fmt"
	"slices"
)

const EOFByte byte = byte(3)    // ETX (End of Text), see ASCII table
var EOFString = string(EOFByte) // "\x03"

var symbolsMap = map[string]TokenType{
	// Operators
	"=":  "ASSIGN",
	"+":  "PLUS",
	"-":  "MINUS",
	"!":  "BANG",
	"*":  "ASTERISK",
	"/":  "SLASH",
	"<":  "LT",
	">":  "GT",
	"==": "EQ",
	"!=": "NOT_EQ",
	// Delimiters
	EOFString: "EOF",
	",":       "COMMA",
	";":       "SEMICOLON",
	"(":       "LPAREN",
	")":       "RPAREN",
	"{":       "LBRACE",
	"}":       "RBRACE",
}

var SymbolsList = []byte{
	'=', '+', '-', '*', '/', '<', '>', '!', ',', ';', '(', ')', '{', '}',
}

func IsSymbol(char byte) bool {
	return slices.Contains(SymbolsList, char)
}

func IsSymbolRec(char byte) bool {
	// how to define recursive nested function in Go
	var IsSymbolRecInner func(byte, int, []byte) bool                    // signature first
	IsSymbolRecInner = func(char byte, position int, list []byte) bool { // then inline function definition
		if position >= len(list) { // empty / exhausted list case
			return false
		} else if list[position] == char {
			return true
		} else {
			return IsSymbolRecInner(char, position+1, list)
		}
	}
	return IsSymbolRecInner(char, 0, SymbolsList)
}

func GetSymbol(str string) (Token, error) {
	if name, ok := symbolsMap[str]; ok {
		return Token{Type: name, Literal: str}, nil
	}
	return Token{}, fmt.Errorf("symbolsMap[str]: %w is not in the map", errors.New(str))
}

func GetEOFToken() (Token, error) {
	return GetSymbol(EOFString)
}

// go-like enum
// these strings will be resolved as TokenType
// const (
// 	ILLEGAL = "ILLEGAL"
// 	// EOF     = "EOF"
// 	EOF = "\x03"
// 	// Identifiers + literals
// 	IDENT = "IDENT" // variable names: add, foobar, x, y, ...
// 	INT   = "INT"   // 1343456
// 	// Operators
// 	ASSIGN   = "="
// 	PLUS     = "+"
// 	MINUS    = "-"
// 	BANG     = "!"
// 	ASTERISK = "*"
// 	SLASH    = "/"
// 	LT       = "<"
// 	GT       = ">"
// 	EQ       = "=="
// 	NOT_EQ   = "!="
// 	// Delimiters
// 	COMMA     = ","
// 	SEMICOLON = ";"
// 	LPAREN    = "("
// 	RPAREN    = ")"
// 	LBRACE    = "{"
// 	RBRACE    = "}"
// 	// Keywords
// 	FUNCTION = "FUNCTION"
// 	LET      = "LET"
// 	TRUE     = "TRUE"
// 	FALSE    = "FALSE"
// 	IF       = "IF"
// 	ELSE     = "ELSE"
// 	RETURN   = "RETURN"
// )

// arbitrary choice for the EOFByte
// const EOFByte byte = byte(0) // corresponding string is "\x00"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// var keywords = map[string]TokenType{
// 	"fn":  FUNCTION,
// 	"let": LET,
// 	// "true":   TRUE,
// 	// "false":  FALSE,
// 	// "if":     IF,
// 	// "else":   ELSE,
// 	// "return": RETURN,
// }

// func LookupIdent(ident string) TokenType {
// 	if tok, ok := keywords[ident]; ok {
// 		return tok
// 	}
// 	return IDENT
// }

// NOOOOOOOOOOOOOOOOO
// type TokenMap map[TokenType]string

// func GetTokenMap() TokenMap {
// 	return TokenMap{
// 		// Identifiers + literals
// 		"ILLEGAL": "ILLEGAL",
// 		"EOF":     "EOF",
// 		// Operators
// 		"IDENT":    "IDENT",
// 		"INT":      "INT",
// 		"ASSIGN":   "=",
// 		"PLUS":     "+",
// 		"MINUS":    "-",
// 		"BANG":     "!",
// 		"ASTERISK": "*",
// 		"SLASH":    "/",
// 		"LT":       "<",
// 		"GT":       ">",
// 		"EQ":       "==",
// 		"NOT_EQ":   "!=",
// 		// Delimiters
// 		"COMMA":     ",",
// 		"SEMICOLON": ";",
// 		"LPAREN":    "(",
// 		"RPAREN":    ")",
// 		"LBRACE":    "{",
// 		"RBRACE":    "}",
// 		// Keywords
// 		"FUNCTION": "FUNCTION",
// 		"LET":      "LET",
// 		"TRUE":     "TRUE",
// 		"FALSE":    "FALSE",
// 		"IF":       "IF",
// 		"ELSE":     "ELSE",
// 		"RETURN":   "RETURN",
// 	}
// }
