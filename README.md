# Termina (name pending)

Termina (pending) is a programming language, but it was designed and made by me, instead of someone else.

## Disclaimer

- This doesn't exist in any implementation presently. I plan to start working on a compiler Soon:tm:
- I have no idea how feasible implementing any of these things are going to be for me. The end result could be drastically different, but this is what I kinda think I'd like a language to look like
- Everything here is subject to change any amount at any time right now
- No assurance that ideas are consistently expressed throughout this doc. It's very free form thought. I might miss something in an example, or contradict a previous statement in a later block. I might say the same thing in 4 different places. 

## Why?

Cuz it sounds really fun! I like learning things, and this seems like a really fun thing to learn. I've been using coding languages regularly for almost 10 years at this point, and I would like to have a better grasp on what is actually happening. I know the buzzwords like "garbage collector" and "bytecode" and "segmentation fault", but... what's actually happening? I dunno, but I sure would like to!

Also, I really like making cool shit, and being able to say "I made a programming language" sounds like some really cool shit.

## Design Goals

Really high level right now. I'm going to probably start with a bunch of indiviual ideas for what I'd like stuff to look like, and then I'm going to have to look over it as a whole and figure out how to tie all of those pieces together, and make any necessary sacrifices.

I'm also just writing stuff out as I think of it, without any regard for difficulty or where it would lie in the roadmap. Once I feel like I've got enough stuff sorted out and a structure to go off of, I'll group stuff out in "waves"

### Things I Know I Want (But Might Find Reasons Against)

- Static typing
- Object oriented, but IDK to what extent yet
- Classes and abstract classes with inheritance
- Loops (for, foreach, while)
- Comments start with `#`
- Multiple value returns from a function
- Allow both single and double quotes
- Easy "this or that" conditional assignment (ternary)
- Try to limit colons outside of actual `this: that` usage. Stuff like `std::fmt::doThing` looks SO ugly

### Things That Sound Neat But Don't Need If Hard

- Implicitly infer function scope based on function naming scheme. It's both a keyword saver and a formatting enforcer. Go does this and I'm a huge fan.
  - `func PublicMethod()`
  - `func privateMethod()` (maybe prefix with underscore, but I like uppercase vs lowercase)
- File I/O of any sort
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
  - Doesn't need to be I/O, could be used for anything
  - Also would have the option to defer a call normall just by doing `defer thing.Func()` or `defer DoFunc(thing)`
- Modules / packages / dependencies / refer to other file's functions somehow

### Things I'm Iffy On

- Semicolons as line endings
- Interfaces (don't hate me, I just don't understand them well enough)
- Anonymous functions / lambdas (see above line)

### Things I Actively Don't Want

- Using two symbols to assign a variable, I hate typing `someVar := 5` (he says, while likely writing this compiler in Go)
- Really just two operators next to each other for anything that aren't combining the two operators
  - `<=` is totally fine, cuz it's the combo of `<` and `=`
  - `->` and `:=` are stupid and I feel like there are almost always better options available
  - `==` kinda breaks this rule, but it's too pervasive at this point, and the thought of single equals doing both assignment and comparison just feels wrong to me, so I'm gonna let it slide.... But it better stay on its best behavior


## References

- [Writing an Interpreter in Go](https://interpreterbook.com)
- [Compiler in Python](https://medium.com/@marcelogdeandrade/writing-your-own-programming-language-and-compiler-with-python-a468970ae6df)
- [TeenyTiny Compiler (also python)](https://austinhenley.com/blog/teenytinycompiler1.html)
- [Pinecone Development](https://www.freecodecamp.org/news/the-programming-language-pipeline-91d3f449c919/)
- [Crafting Interpreters Book](https://craftinginterpreters.com/)
