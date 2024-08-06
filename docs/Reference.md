# Reference

I doubt this document is going to adhere to any standard, but I'm just documenting things about the language that I think need to be formalized in some way.

## Language Reservations

### Keywords

- return
  - Indicates that the remainder of the lineis to be considered the evaluation of the containing function
- var
  - Indicates an assignment
- func
  - Starts a function declaration
- class
  - Starts a class declaration
- interface
  - Starts an interface declaration

### Data Types

- integers --- (int)
- floats ----- (float)
- strings ---- (str)
- characters - (char)
- booleans --- (bool)

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
