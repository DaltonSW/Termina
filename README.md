# Termina (pending)

Termina (pending) is a programming language, but it was made by me instead of someone else. It is, first and foremost, a learning opportunity, and something that I would love to be able to say that I accomplished (whatever I define that to mean).

## Why?

Cuz I wanna. It sounds neat and fun.

## Design Goals

Really high level right now. I'm going to probably start with a bunch of indiviual ideas for what I'd like stuff to look like, and then I'm going to have to look over it as a whole and figure out how to tie all of those pieces together, and make any necessary sacrifices.

I'm also just writing stuff out as I think of it, without any regard for difficulty or where it would lie in the roadmap. Once I feel like I've got enough stuff sorted out and a structure to go off of, I'll group stuff out in "waves"

### Things I Know I Want (But Might Find Reasons Against)

- Static typing
- Object oriented, but IDK to what extent yet
- Classes and abstract classes with inheritance
- Loops (for, foreach, while)
- Comments start with `#`
- Multiple returns from a function
- Allow both single and double quotes
- Ternaries/simple conditional assignment
- Try to limit colons outside of actual `this: that` usage. Stuff like `std::fmt::doThing` looks SO ugly

### Things That Sound Neat But Don't Need If Hard

- Implicitly infer function scope based on function naming scheme. It's both a keyword saver and a formatting enforcer
  - `func PublicMethod()`
  - `func privateMethod()` (maybe prefix with underscore, but I like uppercase vs lowercase)
- Some concept of an auto-deferable function, which essentially acts like an "end of current function" callback
  ```
    func OpenFile(str filename) io.File {
        var f = os.Open(filename)
    } (CloseFile(f))

    func CloseFile(io.File file) {
        os.Close(file)
    }
    ```
  - And then you could call `var file = defer OpenFile("filename.txt")`, which would automatically call `CloseFile(file)` before the function returns execution to its caller
    - It's like Go's system, but... nicer
- Modules / packages / dependencies / refer to other file's functions somehow

### Things I'm Iffy On

- Semicolons as line endings
- Interfaces (don't hate me, I just don't understand them)
- Anonymous functions / lambdas (see above line)

### Things I Actively Don't Want

- Using two symbols to assign a variable, I hate typing `someVar := 5` (he says, while likely writing this compiler in Go)
- Fat arrow syntax. Really just two operators next to each other for anything that aren't combining the two operators
  - `<=` is totally fine, cuz it's the combo of `<` and `=`
  - `->` and `:=` are stupid and I feel like there are almost always better options available
  - `==` kinda breaks this rule, but it's too pervasive at this point, and the thought of single equals doing both assignment and comparison just feels wrong to me, so I'm gonna let it slide.... But it better stay on its best behavior

## Syntax Things

### Examples

### Classes

### Functions

#### Declarations

`func FuncName(type paramName, ...) returnType {} `

## Reserved/Special Things

### Symbols

- LPAREN
- RPAREN
- LBRACK
- RBRACK
- LBRACE
- RBRACE
- PLUS
- MINUS
- STAR
- SLASH
- BSLASH
- QMARK
- EXCLAM
- SQUOTE
- DQUOTE
- COLON
- SEMIC
- COMMA
- PERIOD
- EQUALS
- GREATER
- LESS
- PIPE
- AMPER

(Do we need to 'reserve' a hashtag, or do we just entirely yeet the comments during Step 1/lexing?)

### Keywords

- `class`
- `return`
- `func`
- `defer`
- `interface` (?)
- `abstract` (?)
- `import` (?)

#### Types

- `str`
- `char`
- `int`
- `float`
- `bool`
- `null`

(if differentiating between `int`/`float` and `str`/`char` becomes too difficult, I'll collapse them into `num` and `str`, where a `char` is just a `str` with length 1)

## References

- [Compiler in Python](https://medium.com/@marcelogdeandrade/writing-your-own-programming-language-and-compiler-with-python-a468970ae6df)
- [TeenyTiny Compiler (also python)](https://austinhenley.com/blog/teenytinycompiler1.html)
- [Pinecone Development](https://www.freecodecamp.org/news/the-programming-language-pipeline-91d3f449c919/)
- [Crafting Interpreters Book](https://craftinginterpreters.com/)
