package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	char         byte // current char under examination
	position     int  // current position in input (points to current char)
	nextPosition int  // current reading position in input (after current char)
}

func Do(input string, position int) (token.Token, error) {
	if position+1 > len(input) {
		return token.Token{Type: "EOF", Literal: token.EOFString}, nil
	}
	char := input[position]
	switch {
	case token.IsSymbol(char):
		return token.GetSymbol(string(char))
	default:
		return token.GetEOFToken()
	}
}

// func NextToken(l Lexer) (Lexer, token.Token) {
// 	// l := ReadChar(l1)

// 	// l.skipWhitespace()

// 	switch l.char {
// 	case '=':
// 		return l, newToken(token.ASSIGN, l.char)
// 		// if l.peekChar() == '=' {
// 		// 	ch := l.char
// 		// 	l.readChar()
// 		// 	literal := string(ch) + string(l.char)
// 		// 	return token.Token{Type: token.EQ, Literal: literal}
// 		// } else {
// 		// 	return newToken(token.ASSIGN, l.char)
// 		// }
// 	case '+':
// 		return l, newToken(token.PLUS, l.char)
// 	case '-':
// 		return l, newToken(token.MINUS, l.char)
// 	// case '!':
// 	// 	if l.peekChar() == '=' {
// 	// 		ch := l.char
// 	// 		l.readChar()
// 	// 		literal := string(ch) + string(l.char)
// 	// 		return token.Token{Type: token.NOT_EQ, Literal: literal}
// 	// 	} else {
// 	// 		return newToken(token.BANG, l.char)
// 	// 	}
// 	case '/':
// 		return l, newToken(token.SLASH, l.char)
// 	case '*':
// 		return l, newToken(token.ASTERISK, l.char)
// 	case '<':
// 		return l, newToken(token.LT, l.char)
// 	case '>':
// 		return l, newToken(token.GT, l.char)
// 	case ';':
// 		return l, newToken(token.SEMICOLON, l.char)
// 	case ',':
// 		return l, newToken(token.COMMA, l.char)
// 	case '{':
// 		return l, newToken(token.LBRACE, l.char)
// 	case '}':
// 		return l, newToken(token.RBRACE, l.char)
// 	case '(':
// 		return l, newToken(token.LPAREN, l.char)
// 	case ')':
// 		return l, newToken(token.RPAREN, l.char)
// 	case 0: // NULL char at the end of file
// 		return l, newToken(token.EOF, token.EOFByte)
// 	default: // not a symbol
// 		if isLetter(l.char) {
// 			l, identifier := readIdentifier(l)
// 			return l, token.Token{Type: token.LookupIdent(identifier), Literal: identifier}
// 		} else { // we cannot handle what the current character is
// 			// return newToken(token.EOF, token.EOFByte)
// 			return l, newToken(token.ILLEGAL, l.char)

// 		}
// 		// var tok token.Token
// 		// tok.Literal = ""
// 		// tok.Type = token.EOF
// 		// return tok
// 		// default:
// 		// 	if isLetter(l.char) {
// 		// 		tok.Literal = l.readIdentifier()
// 		// 		tok.Type = token.LookupIdent(tok.Literal)
// 		// 		return tok
// 		// 	} else if isDigit(l.char) {
// 		// 		tok.Type = token.INT
// 		// 		tok.Literal = l.readNumber()
// 		// 		return tok
// 		// 	} else {
// 		// 		return newToken(token.ILLEGAL, l.char)
// 		// 	}
// 	}

// 	// l.readChar()
// 	// return tok
// }

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// func (l *Lexer) readChar() {
// 	if l.nextPosition >= len(l.input) {
// 		l.char = 0
// 	} else {
// 		l.char = l.input[l.nextPosition]
// 	}
// 	l.position = l.nextPosition
// 	l.nextPosition += 1
// }

func ReadChar(lexer Lexer) Lexer {
	if lexer.nextPosition >= len(lexer.input) {
		return Lexer{
			input: lexer.input,
			char:  token.EOFByte,
			// position:     0,
			// nextPosition: 0,
		}
	} else {
		return Lexer{
			input:        lexer.input,
			char:         lexer.input[lexer.nextPosition],
			position:     lexer.nextPosition,
			nextPosition: lexer.nextPosition + 1,
		}
	}
}

// func (l *Lexer) skipWhitespace() {
// 	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
// 		l.readChar()
// 	}
// }

// func (l *Lexer) peekChar() byte {
// 	if l.nextPosition >= len(l.input) {
// 		return 0
// 	} else {
// 		return l.input[l.nextPosition]
// 	}
// }

func readIdentifier(l Lexer) (Lexer, string) {
	result := []byte{l.char} // the first character was identified as a letter
	nextPosition := l.position + 1
	for {
		nextChar := l.input[nextPosition]
		if isLetter(nextChar) {
			result = append(result, nextChar)
			nextPosition += 1
			l.position += 1
		} else {
			break
		}
	}
	return l, string(result)
}

// func (l *Lexer) readNumber() string {
// 	position := l.position
// 	for isDigit(l.char) {
// 		l.readChar()
// 	}
// 	return l.input[position:l.position]
// }

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// func isDigit(ch byte) bool {
// 	return '0' <= ch && ch <= '9'
// }
