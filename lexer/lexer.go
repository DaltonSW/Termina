package lexer

import "github.com/charmbracelet/log"
import "dalton.dog/termina/token"

type Lexer struct {
	tokens    []token.TokenType
	input     string
	curPos    int
	posToRead int
	curChar   byte
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

	switch lexer.curChar {
	case '*':
		newToken = *token.MakeNewToken(token.STAR, lexer.curChar)
	case '/':
		newToken = *token.MakeNewToken(token.SLASH, lexer.curChar)
	case '_':
		newToken = *token.MakeNewToken(token.UNDRSCOR, lexer.curChar)
	case '-':
		newToken = *token.MakeNewToken(token.MINUS, lexer.curChar)
	case '+':
		newToken = *token.MakeNewToken(token.PLUS, lexer.curChar)
	case '=':
		newToken = *token.MakeNewToken(token.EQUALS, lexer.curChar)
	case '(':
		newToken = *token.MakeNewToken(token.LPAREN, lexer.curChar)
	case ')':
		newToken = *token.MakeNewToken(token.RPAREN, lexer.curChar)
	case '[':
		newToken = *token.MakeNewToken(token.LBRACK, lexer.curChar)
	case ']':
		newToken = *token.MakeNewToken(token.RBRACK, lexer.curChar)
	case '{':
		newToken = *token.MakeNewToken(token.LBRACE, lexer.curChar)
	case '}':
		newToken = *token.MakeNewToken(token.RBRACE, lexer.curChar)
	}

	lexer.readNextByte()

	return newToken
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
