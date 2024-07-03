package lexer

import "github.com/charmbracelet/log"
import "dalton.dog/termina/token"

type Lexer struct {
	tokens     []token.TokenType
	input      string
	curPos     int
	posToRead  int
	curChar    byte
	curLineNum int
}

func MakeNewLexer(textIn string) *Lexer {
	lexer := &Lexer{}
	lexer.input = textIn

	log.Info("Making new lexer", "input", textIn)
	lexer.readNextByte()

	return lexer
}

func (lexer *Lexer) GetNextToken() token.Token {
	var newToken token.Token
	var newTokenType token.TokenType
	var newTokenLiteral string

	lexer.curLineNum = 1

	// Single Char Tokens
	switch lexer.curChar {
	// Operators
	case '+':
		newTokenType = token.PLUS
	case '-':
		newTokenType = token.MINUS
	case '*':
		newTokenType = token.STAR
	case '/':
		newTokenType = token.SLASH
	case '!':
		newTokenType = token.EXCLAM

	// Comparison
	case '<':
		newTokenType = token.LESSTHAN
	case '>':
		newTokenType = token.GRTRTHAN

	// Grouping
	case '(':
		newTokenType = token.LPAREN
	case ')':
		newTokenType = token.RPAREN
	case '[':
		newTokenType = token.LBRACK
	case ']':
		newTokenType = token.RBRACK
	case '{':
		newTokenType = token.LBRACE
	case '}':
		newTokenType = token.RBRACE
	// Delimiters

	case ';':
		newTokenType = token.SEMICOL
	case ':':
		newTokenType = token.COLON
	case '_':
		newTokenType = token.UNDRSCOR
	case '\n':
		newTokenType = token.NEWLINE
		lexer.curLineNum += 1
	// Misc

	case '=':
		newTokenType = token.EQUALS
	case '\\':
		newTokenType = token.BSLASH

	case 0:
		newTokenType = token.EOF

	default:
		newTokenType = token.ILLEGAL
	}

	if newTokenType == token.ILLEGAL {
		if isCharacter(lexer.curChar) {
			newToken = lexer.tryReadWord()
		} else if isDigit(lexer.curChar) {
			newToken = lexer.tryReadLiteral()
		} else {
			newToken = *token.MakeNewToken(token.ILLEGAL, "Illegal", lexer.curLineNum)

		}

	} else {
		newTokenLiteral = string(lexer.curChar)
		newToken = *token.MakeNewToken(newTokenType, newTokenLiteral, lexer.curLineNum)
	}

	lexer.readNextByte()

	return newToken
}

func (lexer *Lexer) tryReadWord() token.Token {
	var newTokenType token.TokenType
	startPos := lexer.curPos

	for isCharacter(lexer.curChar) {
		lexer.readNextByte()
	}

	identString := lexer.input[startPos:lexer.curPos]

	newType, ok := token.Keywords[identString]
	if ok {
		newTokenType = newType
	} else {
		newTokenType = token.IDENT
	}

	newToken := *token.MakeNewToken(newTokenType, identString, lexer.curLineNum)

	return newToken
}

func (lexer *Lexer) tryReadLiteral() token.Token {
	startPos := lexer.curPos

	for isCharacter(lexer.curChar) {
		lexer.readNextByte()
	}

	identString := lexer.input[startPos:lexer.curPos]

	newToken := *token.MakeNewToken(token.IDENT, identString, lexer.curLineNum)

	return newToken
}

func isCharacter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.curChar == ' ' || lexer.curChar == '\t' || lexer.curChar == '\r' {
		lexer.readNextByte()
	}
}

func (lexer *Lexer) readNextByte() {
	if lexer.posToRead >= len(lexer.input) {
		lexer.curChar = 0
	} else {
		lexer.curChar = lexer.input[lexer.posToRead]
	}

	lexer.curPos = lexer.posToRead
	lexer.posToRead += 1
}
