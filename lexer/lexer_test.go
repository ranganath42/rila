package lexer_test

import (
	"testing"

	"github.com/ranganath42/rila/lexer"
	"github.com/ranganath42/rila/token"
)

func TestNewToken(t *testing.T) {
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
		name            string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"let", token.LET, "let"},
		{"five", token.IDENT, "five"},
		{"=", token.ASSIGN, "="},
		{"5", token.INT, "5"},
		{";", token.SEMICOLON, ";"},
		{"let", token.LET, "let"},
		{"ten", token.IDENT, "ten"},
		{"=", token.ASSIGN, "="},
		{"10", token.INT, "10"},
		{";", token.SEMICOLON, ";"},
		{"let", token.LET, "let"},
		{"add", token.IDENT, "add"},
		{"=", token.ASSIGN, "="},
		{"fn", token.FUNCTION, "fn"},
		{"(", token.LPAREN, "("},
		{"x", token.IDENT, "x"},
		{",", token.COMMA, ","},
		{"y", token.IDENT, "y"},
		{")", token.RPAREN, ")"},
		{"{", token.LBRACE, "{"},
		{"x", token.IDENT, "x"},
		{"+", token.PLUS, "+"},
		{"y", token.IDENT, "y"},
		{";", token.SEMICOLON, ";"},
		{"}", token.RBRACE, "}"},
		{";", token.SEMICOLON, ";"},
		{"let", token.LET, "let"},
		{"result", token.IDENT, "result"},
		{"=", token.ASSIGN, "="},
		{"add", token.IDENT, "add"},
		{"(", token.LPAREN, "("},
		{"five", token.IDENT, "five"},
		{",", token.COMMA, ","},
		{"ten", token.IDENT, "ten"},
		{")", token.RPAREN, ")"},
		{";", token.SEMICOLON, ";"},
		{"!", token.BANG, "!"},
		{"-", token.MINUS, "-"},
		{"/", token.SLASH, "/"},
		{"*", token.ASTERISK, "*"},
		{"5", token.INT, "5"},
		{";", token.SEMICOLON, ";"},
		{"5", token.INT, "5"},
		{"<", token.LT, "<"},
		{"10", token.INT, "10"},
		{">", token.GT, ">"},
		{"5", token.INT, "5"},
		{";", token.SEMICOLON, ";"},
		{"if", token.IF, "if"},
		{"(", token.LPAREN, "("},
		{"5", token.INT, "5"},
		{"<", token.LT, "<"},
		{"10", token.INT, "10"},
		{")", token.RPAREN, ")"},
		{"{", token.LBRACE, "{"},
		{"return", token.RETURN, "return"},
		{"true", token.TRUE, "true"},
		{";", token.SEMICOLON, ";"},
		{"}", token.RBRACE, "}"},
		{"else", token.ELSE, "else"},
		{"{", token.LBRACE, "{"},
		{"return", token.RETURN, "return"},
		{"false", token.FALSE, "false"},
		{";", token.SEMICOLON, ";"},
		{"}", token.RBRACE, "}"},
		{"10", token.INT, "10"},
		{"==", token.EQ, "=="},
		{"10", token.INT, "10"},
		{";", token.SEMICOLON, ";"},
		{"10", token.INT, "10"},
		{"!=", token.NOT_EQ, "!="},
		{"9", token.INT, "9"},
		{";", token.SEMICOLON, ";"},

		{"EOF", token.EOF, ""},
	}

	lex := lexer.New(input)
	for _, tt := range tests {
		tok := lex.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("[%s] invalid token type. expected: %q, got: %q", tt.name, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("[%s] invalid token literal. expected: %q, got: %q", tt.name, tt.expectedLiteral, tok.Literal)
		}
	}
}
