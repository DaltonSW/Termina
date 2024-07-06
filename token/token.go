package token

type TokenType int

const (
	EOF TokenType = iota

	// Operators
	PLUS
	MINUS
	SLASH
	STAR
	EXCLAM

	// Comparison
	GRTRTHAN
	LESSTHAN
	GRTREQUAL
	LESSEQUAL
	ISEQUAL
	NOTEQUAL
	IF
	ELSE

	// Grouping
	LPAREN
	RPAREN
	LBRACK
	RBRACK
	LBRACE
	RBRACE

	// Delimiters
	SEMICOL
	COLON
	UNDRSCOR
	NEWLINE
	PERIOD
	COMMA

	// Keywords
	VAR
	RETURN
	FUNC

	// Identifers & Types
	IDENT
	TYPEDECL

	INTEGER
	FLOAT
	STRING
	CHAR
	BOOLEAN

	// Misc
	EQUALS
	ILLEGAL
	BSLASH
)

var Keywords = map[string]TokenType{
	"func":   FUNC,
	"return": RETURN,
	"var":    VAR,
	"if":     IF,
	"else":   ELSE,
}

var Types = []string{
	"int", "float", "char", "str", "bool",
}

type Token struct {
	Type    TokenType
	Literal string
	LineNum int
}

func MakeNewToken(tokenType TokenType, tokenLiteral string, lineNum int) *Token {
	newToken := &Token{
		Type:    tokenType,
		Literal: string(tokenLiteral),
		LineNum: lineNum,
	}

	return newToken
}
