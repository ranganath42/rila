package lexer

import "github.com/ranganath42/rila/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()
	return &l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	twoCharTokens := map[byte]token.Token{
		'=': token.NewToken(token.EQ, "=="),
		'!': token.NewToken(token.NOT_EQ, "!="),
		'<': token.NewToken(token.LTE, "<="),
		'>': token.NewToken(token.GTE, ">="),
	}
	if tok, ok := twoCharTokens[l.ch]; ok {
		if l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			return tok
		}
	}

	oneCharTokens := map[byte]token.Token{
		'+': token.NewToken(token.PLUS, string(l.ch)),
		',': token.NewToken(token.COMMA, string(l.ch)),
		';': token.NewToken(token.SEMICOLON, string(l.ch)),
		'(': token.NewToken(token.LPAREN, string(l.ch)),
		')': token.NewToken(token.RPAREN, string(l.ch)),
		'{': token.NewToken(token.LBRACE, string(l.ch)),
		'}': token.NewToken(token.RBRACE, string(l.ch)),
		'-': token.NewToken(token.MINUS, string(l.ch)),
		'*': token.NewToken(token.ASTERISK, string(l.ch)),
		'/': token.NewToken(token.SLASH, string(l.ch)),
		'<': token.NewToken(token.LT, string(l.ch)),
		'>': token.NewToken(token.GT, string(l.ch)),
		'=': token.NewToken(token.ASSIGN, string(l.ch)),
		'!': token.NewToken(token.BANG, string(l.ch)),
		0:   token.NewToken(token.EOF, ""),
	}
	if tok, ok := oneCharTokens[l.ch]; ok {
		l.readChar()
		return tok
	}

	var tok token.Token
	if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	}
	if isDigit(l.ch) {
		tok.Literal = l.readNumber()
		tok.Type = token.INT
		return tok
	}
	tok = token.NewToken(token.ILLEGAL, string(l.ch))
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter returns true if the given byte is a letter or underscore.
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// isDigit returns true if the given byte is a digit.
// Rila only supports integers, so we don't need to worry about floats or other numeric types
// such as hex, octal, or binary.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
