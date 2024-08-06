# Reference

## Language Reservations

### Keywords

- `return`
  - Indicates that the remainder of the lineis to be considered the evaluation of the containing function
- `var`
  - Indicates an assignment
- `func`
  - Starts a function declaration
- `class`
  - Starts a class declaration
- `interface`
  - Starts an interface declaration
- `defer`
  - Indicates that a function call will happen before the containing function returns
- `import`
  - Starts a block of module import calls
- `const`
  - Declare a constant
- `static`
  - Declares a mutable value that is shared across all implementations of a class
- `abstract`
  - Indicates that a class is abstract. Can't have implementation
- `parent`
  - References the parent class
- `if`, `else`, `elif`
  - Traditional conditional branches
- `switch`, `case`
  - Switch statement conditional branching

### Data Types

#### Literals

- integers --- (int)
- floats ----- (float)
- strings ---- (str)
- characters - (char)
- booleans --- (bool)
- null (is this a type?)

#### Collections

- list
- dict

### Symbols

- `( ) [ ] { } ' "`
  - Groupings 
- `+ - * / %`
  - Math
- `= < > ? & |`
  - Comparison
- `\ ! : ; , .`
  - Misc

## Conditional Evaluation

### Value Literals as Booleans

This is how values of these literal types will be evaluated if used in (as?) a boolean expression.

| Type      | TRUE      | FALSE      |
| ----      | ----      | -----      |
| int (n)   | n != 0    | n == 0     |
| float (f) | f != 0.0  | f == 0.0   |
| str (s)   | s != ""   | s == ""    |
| char (c)  | c != ''   | c == ''    |
| bool (b)  | b == TRUE | b == FALSE |
