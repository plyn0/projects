package lexer

import (
	"interpreter/token"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtendedSnippet(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);

	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral token.Literal
	}{
		{"LET", "let"},
		{"IDENT", "five"},
		{"ASSIGN", "="},
		{"INT", "5"},
		{"SEMICOLON", ";"},
		{"LET", "let"},
		{"IDENT", "ten"},
		{"ASSIGN", "="},
		{"INT", "10"},
		{"SEMICOLON", ";"},
		{"LET", "let"},
		{"IDENT", "add"},
		{"ASSIGN", "="},
		{"FUNCTION", "fn"},
		{"LPAREN", "("},
		{"IDENT", "x"},
		{"COMMA", ","},
		{"IDENT", "y"},
		{"RPAREN", ")"},
		{"LBRACE", "{"},
		{"IDENT", "x"},
		{"PLUS", "+"},
		{"IDENT", "y"},
		{"SEMICOLON", ";"},
		{"RBRACE", "}"},
		{"SEMICOLON", ";"},
		{"LET", "let"},
		{"IDENT", "result"},
		{"ASSIGN", "="},
		{"IDENT", "add"},
		{"LPAREN", "("},
		{"IDENT", "five"},
		{"COMMA", ","},
		{"IDENT", "ten"},
		{"RPAREN", ")"},
		{"SEMICOLON", ";"},
		{"BANG", "!"},
		{"MINUS", "-"},
		{"SLASH", "/"},
		{"ASTERISK", "*"},
		{"INT", "5"},
		{"SEMICOLON", ";"},
		{"INT", "5"},
		{"LT", "<"},
		{"INT", "10"},
		{"GT", ">"},
		{"INT", "5"},
		{"SEMICOLON", ";"},
		{"IF", "if"},
		{"LPAREN", "("},
		{"INT", "5"},
		{"LT", "<"},
		{"INT", "10"},
		{"RPAREN", ")"},
		{"LBRACE", "{"},
		{"RETURN", "return"},
		{"TRUE", "true"},
		{"SEMICOLON", ";"},
		{"RBRACE", "}"},
		{"ELSE", "else"},
		{"LBRACE", "{"},
		{"RETURN", "return"},
		{"FALSE", "false"},
		{"SEMICOLON", ";"},
		{"RBRACE", "}"},
		{"INT", "10"},
		{"EQ", "=="},
		{"INT", "10"},
		{"SEMICOLON", ";"},
		{"INT", "10"},
		{"NOT_EQ", "!="},
		{"INT", "9"},
		{"SEMICOLON", ";"},
		// {"EOF", ""},
		{"EOF", "\x03"},
	}

	position := 0
	for _, value := range tests {

		position = SkipWhitespace(input, position)
		tok, err := Do(input, position)
		if err != nil {
			log.Fatalln(err)
		}
		if tok.Literal == token.EOFLiteral {
			return
		}
		position += len(tok.Literal) // advance the position depending on the length of the token

		require.Equal(t, value.expectedType, tok.Type)
		require.Equal(t, value.expectedLiteral, tok.Literal)
	}
}
func TestSnippet(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral token.Literal
	}{
		{"LET", "let"},
		{"IDENT", "five"},
		{"ASSIGN", "="},
		{"INT", "5"},
		{"SEMICOLON", ";"},
		{"LET", "let"},
		{"IDENT", "ten"},
		{"ASSIGN", "="},
		{"INT", "10"},
		{"SEMICOLON", ";"},
		{"LET", "let"},
		{"IDENT", "add"},
		{"ASSIGN", "="},
		{"FUNCTION", "fn"},
		{"LPAREN", "("},
		{"IDENT", "x"},
		{"COMMA", ","},
		{"IDENT", "y"},
		{"RPAREN", ")"},
		{"LBRACE", "{"},
		{"IDENT", "x"},
		{"PLUS", "+"},
		{"IDENT", "y"},
		{"SEMICOLON", ";"},
		{"RBRACE", "}"},
		{"SEMICOLON", ";"},
		{"LET", "let"},
		{"IDENT", "result"},
		{"ASSIGN", "="},
		{"IDENT", "add"},
		{"LPAREN", "("},
		{"IDENT", "five"},
		{"COMMA", ","},
		{"IDENT", "ten"},
		{"RPAREN", ")"},
		{"SEMICOLON", ";"},
		{"EOF", "\x03"},
	}

	position := 0
	for _, value := range tests {
		position = SkipWhitespace(input, position)
		tok, err := Do(input, position)
		if err != nil {
			log.Fatalln(err)
		}
		if tok.Literal == token.EOFLiteral {
			return
		}
		position += len(tok.Literal) // advance the position depending on the length of the token

		require.Equal(t, value.expectedType, tok.Type)
		require.Equal(t, value.expectedLiteral, tok.Literal)
	}
}

func TestSymbols(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral token.Literal
	}{
		{"ASSIGN", "="},
		{"PLUS", "+"},
		{"LPAREN", "("},
		{"RPAREN", ")"},
		{"LBRACE", "{"},
		{"RBRACE", "}"},
		{"COMMA", ","},
		{"SEMICOLON", ";"},
		{"EOF", "\x03"},
	}

	position := 0
	for _, value := range tests {
		position = SkipWhitespace(input, position)
		tok, err := Do(input, position)
		if err != nil {
			log.Fatalln(err)
		}
		if tok.Literal == token.EOFLiteral {
			return
		}
		position += len(tok.Literal) // advance the position depending on the length of the token

		require.Equal(t, value.expectedType, tok.Type)
		require.Equal(t, value.expectedLiteral, tok.Literal)
	}
}
