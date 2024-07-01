package lexer

import (
	"dalton.dog/termina/token"
	"testing"
)

func TestTokenParsingOperators(t *testing.T) {
	input := "*_-+=/"
	lex := MakeNewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STAR, "*"},
		{token.UNDRSCOR, "_"},
		{token.MINUS, "-"},
		{token.PLUS, "+"},
		{token.EQUALS, "="},
		{token.SLASH, "/"},
	}

	for i, expected := range tests {
		nextToken := lex.GetNextToken()

		if nextToken.Type != expected.expectedType {
			t.Fatalf("tests[%d] failed - Wrong TokenType. Expected %q, got %q", i, expected.expectedType, nextToken.Type)
		}

		if nextToken.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] failed - Wrong TokenLiteral . Expected %q, got %q", i, expected.expectedLiteral, nextToken.Literal)
		}

	}
}
