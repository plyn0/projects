package lexer

import (
	"log"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	// input := `let five = 5;
	// let ten = 10;

	// let add = fn(x, y) {
	//   x + y;
	// };

	// let result = add(five, ten);
	// `

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
		expectedLiteral string
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

		// {token.LET, "let"},
		// {token.IDENT, "five"},
		// {token.ASSIGN, "="},
		// {token.INT, "5"},
		// {token.SEMICOLON, ";"},
		// {token.LET, "let"},
		// {token.IDENT, "ten"},
		// {token.ASSIGN, "="},
		// {token.INT, "10"},
		// {token.SEMICOLON, ";"},
		// {token.LET, "let"},
		// {token.IDENT, "add"},
		// {token.ASSIGN, "="},
		// {token.FUNCTION, "fn"},
		// {token.LPAREN, "("},
		// {token.IDENT, "x"},
		// {token.COMMA, ","},
		// {token.IDENT, "y"},
		// {token.RPAREN, ")"},
		// {token.LBRACE, "{"},
		// {token.IDENT, "x"},
		// {token.PLUS, "+"},
		// {token.IDENT, "y"},
		// {token.SEMICOLON, ";"},
		// {token.RBRACE, "}"},
		// {token.SEMICOLON, ";"},
		// {token.LET, "let"},
		// {token.IDENT, "result"},
		// {token.ASSIGN, "="},
		// {token.IDENT, "add"},
		// {token.LPAREN, "("},
		// {token.IDENT, "five"},
		// {token.COMMA, ","},
		// {token.IDENT, "ten"},
		// {token.RPAREN, ")"},
		// {token.SEMICOLON, ";"},

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
	for i, value := range tests {
		tok, err := Do(input, position)
		if err != nil {
			log.Fatalln(err)
		}
		if tok.Literal == token.EOFString {
			return
		}
		position += 1

		// fmt.Println("===")
		// fmt.Println(string(l.char), i, tok, l.position, 1)
		// l = ReadChar(l)
		// fmt.Println(string(l.char), i, tok, l.position, 2)
		// // tok := l.Nexvalueoken()

		// // update lexer and token
		// l, tok = Nexvalueoken(l)
		// fmt.Println(string(l.char), i, tok, l.position, 3)

		if tok.Type != value.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, value.expectedType, tok.Type)
		}

		if tok.Literal != value.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, value.expectedLiteral, tok.Literal)
		}
	}
}

func TestSymbols(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
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
	for i, value := range tests {
		tok, err := Do(input, position)
		if err != nil {
			log.Fatalln(err)
		}
		if tok.Literal == token.EOFString {
			return
		}
		position += 1

		if tok.Type != value.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, value.expectedType, tok.Type)
		}

		if tok.Literal != value.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, value.expectedLiteral, tok.Literal)
		}
	}
}
