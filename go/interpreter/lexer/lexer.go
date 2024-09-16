package lexer

import (
	"fmt"
	"interpreter/token"
)

func Do(input string, position int) (token.Token, error) {
	if position >= len(input) { // in the entire program, the EOF is created here
		return token.Token{Literal: token.EOFLiteral, Type: "EOF"}, nil
	}
	char := input[position]
	switch {
	case token.IsSymbol(char):
		return token.GetSymbol(token.Literal(char))
	case IsLetter(char):
		ident := GetIdentifier(input, position)
		tokenType := token.LookupIdent(ident)
		return token.Token{Literal: ident, Type: tokenType}, nil
	case IsDigit(char):
		number := GetNumber(input, position)
		return token.Token{Literal: number, Type: "INT"}, nil
	default:
		return token.Token{}, fmt.Errorf("switch on token: char '%s' is not recognized", string(char))
	}
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func GetNumber(input string, position int) token.Literal {
	return getTokenLiteral(input, position, IsDigit)
}

func GetIdentifier(input string, position int) token.Literal {
	return getTokenLiteral(input, position, IsLetter)
}

func getTokenLiteral(input string, position int, predicate func(byte) bool) token.Literal {
	result := []byte{input[position]} // add the first character identified by the predicate function
	for {
		position += 1
		if position > len(input) {
			return token.Literal(result)
		}
		nextChar := input[position]
		if predicate(nextChar) {
			result = append(result, nextChar)
		} else {
			return token.Literal(result)
		}
	}
}

func SkipWhitespace(input string, position int) int {
	if position >= len(input) { // do nothing if end of input is reached
		return position
	}
	for {
		char := input[position]
		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			position += 1
		} else {
			return position
		}
	}
}

// func (l *Lexer) peekChar() byte {
// 	if l.nextPosition >= len(l.input) {
// 		return 0
// 	} else {
// 		return l.input[l.nextPosition]
// 	}
// }
