package lexer

import (
	"pl-i18n/token"
	"strings"
)

// Lexer creates tokens from the input program
type Lexer struct {
	input        string
	position     int // curr position in input (points to curr char)
	readPosition int // current reading position in input (after curr char)
	ch           byte
}

// New creates an instance of Lexer
func New(input string) *Lexer {
	input = replaceKeywords(input)
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func replaceKeywords(input string) string {
	replacer := strings.NewReplacer("정수", "int", "문자", "string", "부동소수", "float", "불", "bool", "널", "NULL", "이거면", "if", "않으면", "else", "펑크션", "func", "산출", "return", "진실", "true", "거짓", "false", "프린트", "print")
	return replacer.Replace(input)
}

// NextToken creates the next token in the program
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '\n':
		tok = newToken(token.NEWLINE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tokLiteral := l.readIdentifier()
			tok.Type = token.LookupIdent(tokLiteral)
			if isIdentType(tok.Type) {
				l.skipWhitespace()
				tok.Literal = l.readIdentifier()
			} else {
				tok.Literal = tokLiteral
			}
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber(&tok)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) readNumber(tok *token.Token) string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if isPeriod(l.ch) {
		l.readChar()
	} else {
		tok.Type = token.INT
		return l.input[position:l.position]
	}
	for isDigit(l.ch) {
		l.readChar()
	}
	tok.Type = token.FLOAT
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isPeriod(ch byte) bool {
	return '.' == ch
}

func isIdentType(ty token.TokenType) bool {
	if ty == token.INTID || ty == token.STRINGID || ty == token.FLOATID || ty == token.BOOLID {
		return true
	}
	return false
}
