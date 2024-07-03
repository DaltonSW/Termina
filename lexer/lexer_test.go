package lexer

import (
	"testing"

	"dalton.dog/termina/token"
	"github.com/charmbracelet/log"
)

func TestTokenParsingOperators(t *testing.T) {
	input := "*_-+=/"
	// log.SetLevel(log.DebugLevel)
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

		if nextToken.Type != expected.expectedType || nextToken.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] failed - Token mismatch.\nExpected %q (%q)\nGot %q (%q)", i, expected.expectedLiteral, expected.expectedType, nextToken.Literal, nextToken.Type)
		}

		log.Info("Token", "Str", nextToken.Literal, "Type", nextToken.Type)
	}
}

func TestTokenParsingGroupers(t *testing.T) {
	input := "()[]{}}][)({"
	// log.SetLevel(log.DebugLevel)
	lex := MakeNewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACK, "]"},
		{token.LBRACK, "["},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.LBRACE, "{"},
	}

	for i, expected := range tests {
		nextToken := lex.GetNextToken()

		if nextToken.Type != expected.expectedType || nextToken.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] failed - Token mismatch.\nExpected %q (%q)\nGot %q (%q)", i, expected.expectedLiteral, expected.expectedType, nextToken.Literal, nextToken.Type)
		}

		log.Info("Token", "Str", nextToken.Literal, "Type", nextToken.Type)
	}
}

func TestTokenParsingNewLines(t *testing.T) {
	input := `




	`

	lex := MakeNewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
	}

	for i, expected := range tests {
		nextToken := lex.GetNextToken()

		if nextToken.Type != expected.expectedType || nextToken.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] failed - Token mismatch.\nExpected %q (%q)\nGot %q (%q)", i, expected.expectedLiteral, expected.expectedType, nextToken.Literal, nextToken.Type)
		}

		log.Info("Token", "Str", nextToken.Literal, "Type", nextToken.Type)
	}
}
func TestTokenParsingNumberLiteral(t *testing.T) {
	input := "5 10 15 20 100 3.14"

	lex := MakeNewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LITERAL, "5"},
		{token.LITERAL, "10"},
		{token.LITERAL, "15"},
		{token.LITERAL, "20"},
		{token.LITERAL, "100"},
		{token.LITERAL, "3.14"},
	}

	for i, expected := range tests {
		nextToken := lex.GetNextToken()

		if nextToken.Type != expected.expectedType || nextToken.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] failed - Token mismatch.\nExpected %q (%q)\nGot %q (%q)", i, expected.expectedLiteral, expected.expectedType, nextToken.Literal, nextToken.Type)
		}

		log.Info("Token", "Str", nextToken.Literal, "Type", nextToken.Type)
	}
}

func TestTokenParsingSimpleProgram(t *testing.T) {
	input := `var x = 5
	var y = 10

	func add(int x, int y) {
		return x + y
	}

	var sum = add(x, y)

	exit(0)
	`
	// log.SetLevel(log.DebugLevel)
	lex := MakeNewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "x"},
		{token.EQUALS, "="},
		{token.LITERAL, "5"},
		{token.NEWLINE, "\n"},

		{token.VAR, "var"},
		{token.IDENT, "y"},
		{token.EQUALS, "="},
		{token.LITERAL, "10"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},

		{token.FUNC, "func"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.INTEGER, "int"},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.INTEGER, "int"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		//		{token.NEWLINE, "nl"},

		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		//		{token.NEWLINE, "nl"},
		{token.RBRACE, "}"},
		//		{token.NEWLINE, "nl"},
		//		{token.NEWLINE, "nl"},

		{token.VAR, "var"},
		{token.IDENT, "sum"},
		{token.EQUALS, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		//		{token.NEWLINE, "nl"},
		//		{token.NEWLINE, "nl"},

		{token.IDENT, "exit"},
		{token.LPAREN, "("},
		{token.LITERAL, "0"},
		{token.RPAREN, ")"},
	}

	for i, expected := range tests {
		nextToken := lex.GetNextToken()

		if nextToken.Type != expected.expectedType || nextToken.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] failed - Token mismatch.\nExpected %q (%q)\nGot %q (%q)", i, expected.expectedLiteral, expected.expectedType, nextToken.Literal, nextToken.Type)
		}

		log.Info("Token", "Str", nextToken.Literal, "Type", nextToken.Type)
	}

}
