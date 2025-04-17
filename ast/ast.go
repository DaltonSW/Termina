// ast represents the Abstract Syntax Tree, and the components that build it up
package ast

import "dalton.dog/termina/token"

/*
Defines 3 different interfaces as building blocks for the AST:
	- Node (a single node of the whole AST)
	- Statement (doesn't return a value)
	- Expression (returns a value)

The AST is going to use 'Program' as the root node, and every Program is going to consist of a slice of valid 'Statements'
*/

// Node is a single node of the entire AST.
type Node interface {
	TokenLiteral() string
}

// Statement is something that does NOT return a value.
type Statement interface {
	Node
	statementNode()
}

// Expression is something that DOES return a value.
type Expression interface {
	Node
	expressionNode()
}

// Program is a Node that represents the root of every AST produced.
// Every valid program is just going to be a sequence of valid Statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Identifier is an Expression that is a representation of A Thing, such as x or valueProducingFunc.
type Identifier struct {
	Token token.Token // IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Let is a Statement that handles association of an Identifier with an Expression, such as `let x = 5`.
type LetStatement struct {
	Token token.Token // LET
	Name  *Identifier
	Value Expression
}
