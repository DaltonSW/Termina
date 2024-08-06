package parser

/*
Defines 3 different interfaces as building blocks for the AST:
	- Node (a single node of the whole AST)
	- Statement (doesn't return a value)
	- Expression (returns a value)

The AST is going to use 'Program' as the root node, and every Program is going to consist of a slice of valid 'Statements'

*/
