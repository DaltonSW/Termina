package token

type TokenType int

const (
	// Operators
	PLUS TokenType = iota
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

	// Keywords
	VAR
	RETURN
	FUNC

	// Identifers & Types
	IDENT
	INTEGER
	FLOAT
	STRING
	CHAR
	BOOLEAN

	// Misc
	EQUALS
	ILLEGAL
	EOF
	BSLASH
)

type Token struct {
	Type    TokenType
	Literal string
}

func MakeNewToken(tokenType TokenType, tokenLiteral byte) *Token {
	newToken := &Token{
		Type:    tokenType,
		Literal: string(tokenLiteral),
	}

	return newToken
}
