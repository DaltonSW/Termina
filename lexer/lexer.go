package lexer

import (
	//	"strings"

	"dalton.dog/termina/token"
	"github.com/charmbracelet/log"
)

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
	//lexer.readNextByte()

	return lexer
}

func (lexer *Lexer) GetNextToken() token.Token {
	var newToken token.Token
	var newTokenType token.TokenType
	var newTokenLiteral string

	lexer.readNextByte()
	lexer.skipWhitespace()

	log.Info("cur char", "Char", lexer.curChar)

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

	// Was able to match a single char. Build new token and return
	if newTokenType != token.ILLEGAL {
		newTokenLiteral = string(lexer.curChar)
		newToken = *token.MakeNewToken(newTokenType, newTokenLiteral, lexer.curLineNum)
		return newToken
	}

	log.Info("Couldn't match on single char.")
	if isCharacter(lexer.curChar) {
		newToken = lexer.tryReadWord()
	} else if isDigit(lexer.curChar) {
		newToken = lexer.tryReadNumber()
	} else {
		newToken = *token.MakeNewToken(token.ILLEGAL, string(lexer.curChar), lexer.curLineNum)

	}

	return newToken
}

func (lexer *Lexer) tryReadWord() token.Token {
	//log.Info("Trying to read word")
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

func (lexer *Lexer) tryReadNumber() token.Token {
	// var newTokenType token.TokenType

	//log.Info("Trying to read number")
	startPos := lexer.curPos

	for isDigit(lexer.curChar) || lexer.curChar == '.' {
		lexer.readNextByte()
	}

	identString := lexer.input[startPos:lexer.curPos]

	// if strings.Contains(identString, ".") {
	// 	newTokenType = token.FLOAT
	// } else {
	// 	newTokenType = token.INTEGER
	// }

	newToken := *token.MakeNewToken(token.LITERAL, identString, lexer.curLineNum)

	return newToken
}

// func (lexer *Lexer) tryReadString() token.Token {
// 	return nil
// }

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
