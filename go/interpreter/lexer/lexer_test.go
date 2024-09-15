package lexer

import (
	"interpreter/token"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnippet(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
`

	// !-/*5;
	// 5 < 10 > 5;

	// if (5 < 10) {
	// 	return true;
	// } else {
	// 	return false;
	// }

	// 10 == 10;
	// 10 != 9;

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

		// {token.BANG, "!"},
		// {token.MINUS, "-"},
		// {token.SLASH, "/"},
		// {token.ASTERISK, "*"},
		// {token.INT, "5"},
		// {token.SEMICOLON, ";"},
		// {token.INT, "5"},
		// {token.LT, "<"},
		// {token.INT, "10"},
		// {token.GT, ">"},
		// {token.INT, "5"},
		// {token.SEMICOLON, ";"},
		// {token.IF, "if"},
		// {token.LPAREN, "("},
		// {token.INT, "5"},
		// {token.LT, "<"},
		// {token.INT, "10"},
		// {token.RPAREN, ")"},
		// {token.LBRACE, "{"},
		// {token.RETURN, "return"},
		// {token.TRUE, "true"},
		// {token.SEMICOLON, ";"},
		// {token.RBRACE, "}"},
		// {token.ELSE, "else"},
		// {token.LBRACE, "{"},
		// {token.RETURN, "return"},
		// {token.FALSE, "false"},
		// {token.SEMICOLON, ";"},
		// {token.RBRACE, "}"},
		// {token.INT, "10"},
		// {token.EQ, "=="},
		// {token.INT, "10"},
		// {token.SEMICOLON, ";"},
		// {token.INT, "10"},
		// {token.NOT_EQ, "!="},
		// {token.INT, "9"},
		// {token.SEMICOLON, ";"},
		// {token.EOF, ""},
	}

	// l := New(input)
	// l := &Lexer{input: input}
	// l.readChar()

	//initialize lexer and token
	// l := Lexer{input: input}
	// var tok token.Token
	// l := ReadChar(l1)

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
		// position += 1
		position += len(tok.Literal) // advance the position depending on the length of the token

		require.Equal(t, value.expectedType, tok.Type, "should be the same")
		require.Equal(t, value.expectedLiteral, tok.Literal, "should be the same")
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
		// position += 1
		position += len(tok.Literal) // advance the position depending on the length of the token

		require.Equal(t, value.expectedType, tok.Type, "should be the same")
		require.Equal(t, value.expectedLiteral, tok.Literal, "should be the same")
		// if tok.Type != value.expectedType {
		// 	t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
		// 		i, value.expectedType, tok.Type)
		// }

		// if tok.Literal != value.expectedLiteral {
		// 	t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
		// 		i, value.expectedLiteral, tok.Literal)
		// }
	}
}

// func TestComputePriceWithVAT(t *testing.T) {
// 	require.Equal(t, computePriceWithVAT(12.4567), 13.14)
// 	require.Equal(t, computePriceWithVAT(1), 1.06) // integer input
// }
