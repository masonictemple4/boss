package lexer

import "github.com/masonictemple4/boss/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (next)
	ch           byte // current char being processed
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// initialize l.ch and l.readPosition
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			currCh := l.ch
			l.readChar()
			lit := string(currCh) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: lit}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			currCh := l.ch
			l.readChar()
			lit := string(currCh) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: lit}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// We exit because readIdentifier() calls readChar() so we don't need to do it again.
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			// exit because readNumber() calls readChar() so we don't have to do it again.
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	// if current character is whitespace read until it isn't.
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// basically an accessor to see what character is at readposition
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// This method needs to read the rest of the identifier/kw
// until it encounters a non-letter.
// It then returns the value of that whole word.
func (l *Lexer) readIdentifier() string {
	currPos := l.position // grab the start before we continue reading.
	// while l.ch is a letter
	for isLetter(l.ch) {
		// read next.
		l.readChar()
	}
	// once we exit this loop l.position will be the last letter in the sequence
	return l.input[currPos:l.position]
}

// CHALLENGE: Enable unicode and utf8 support.
func (l *Lexer) readChar() {
	// Check if we have reached the end of input.
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ascii code for "NUL" or "NULL"
	} else {
		l.ch = l.input[l.readPosition] // read the next character
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// CHALLENGE: Simplify the readNumber/readChar by passing in character identifying functions instead.
func (l *Lexer) readNumber() string {
	currPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[currPos:l.position]
}

// helper/utils
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	// allows for _ in identifier and keywords.
	// CHALLENGE: Add ! and ?
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// CHALLENGE: Floats/Hex notation/octal notation etc..
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
