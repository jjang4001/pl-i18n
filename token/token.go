package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT    = "IDENT"
	INTID    = "INTID"
	STRINGID = "STRINGID"
	FLOATID  = "FLOATID"
	BOOLID   = "BOOLID"

	INT    = "INT"
	STRING = "STRING"
	FLOAT  = "FLOAT"
	BOOL   = "BOOL"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	NEWLINE = "\n"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	WHILE    = "WHILE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"while":  WHILE,
	"return": RETURN,

	// identifiers
	"func":   FUNCTION,
	"int":    INTID,
	"string": STRINGID,
	"float":  FLOATID,
	"bool":   BOOLID,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
