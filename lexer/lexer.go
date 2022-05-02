package lexer

import (
	"github.com/leonardchinonso/leon/token"
	"github.com/leonardchinonso/leon/tools"
)

type Lexer struct {
	in      string
	pos     int
	nextPos int
	ch      byte
}

func NewLexer(in string) *Lexer {
	l := &Lexer{in: in}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.in) {
		l.ch = 0
	} else {
		l.ch = l.in[l.nextPos]
	}
	l.pos = l.nextPos
	l.nextPos++
}

func (l *Lexer) readLiteral(fn func(c byte) bool) string {
	pos := l.pos
	for fn(l.ch) {
		l.readChar()
	}
	return l.in[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.in) {
		return 0
	}
	return l.in[l.nextPos]
}

func (l *Lexer) eatWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) makeTwoCharToken(tt token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	return token.NewTokenFromString(tt, string(ch)+string(l.ch))
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhiteSpace()

	switch l.ch {
	case ';':
		tok = token.NewTokenFromByte(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewTokenFromByte(token.LPAREN, l.ch)
	case ')':
		tok = token.NewTokenFromByte(token.RPAREN, l.ch)
	case '{':
		tok = token.NewTokenFromByte(token.LBRACE, l.ch)
	case '}':
		tok = token.NewTokenFromByte(token.RBRACE, l.ch)
	case ',':
		tok = token.NewTokenFromByte(token.COMMA, l.ch)
	case '+':
		tok = token.NewTokenFromByte(token.PLUS, l.ch)
	case '-':
		tok = token.NewTokenFromByte(token.MINUS, l.ch)
	case '*':
		tok = token.NewTokenFromByte(token.ASTERISK, l.ch)
	case '/':
		tok = token.NewTokenFromByte(token.SLASH, l.ch)
	case '<':
		tok = token.NewTokenFromByte(token.LT, l.ch)
	case '>':
		tok = token.NewTokenFromByte(token.GT, l.ch)
	case '=':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.EQ)
		} else {
			tok = token.NewTokenFromByte(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.makeTwoCharToken(token.NOT_EQ)
		} else {
			tok = token.NewTokenFromByte(token.EXCLAM, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if tools.IsLetter(l.ch) {
			tok.Literal = l.readLiteral(tools.IsLetter)
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if tools.IsDigit(l.ch) {
			tok.Literal = l.readLiteral(tools.IsDigit)
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewTokenFromByte(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
