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

Disclaimer:
- I wrote this out just kinda freeform to see what the syntax would end up like if I just wrote some code snippets how I wanted
- Everything here is subject to change at literally any whim
- Definitely no assurance that some things in these examples don't contradict other parts

`main.trm`
```
# Comments start with an octothorpe

# Main entry point of the program. If you call `term run <file>`, it will run that file's main()
# Needs to return an exit code. 0 is a success, anything greater than that is an error
func main() int {
  print("Hello, mom!");
  defer print("Goodbye, mom!"); # Will run after everything in the function, but before function returns control to caller
  var name = "Dalton"; # String literal
  var month = 12; Number literal
  var day = 25;
  var tempOutside = 79.4; # Float literal
  var tempInside = 69.0; # Local variable naming convention is snakeCase, like with privateFunctions
  printf("Hello {name}, how are you doing on today, {month}/{day}? ", newline=false);
  printf("Did you know that it is {tempOutside} degrees outside?");
  printf("Thankfully it's a nice {tempInside} in here!");

  return 0;
}
```

`shapes.trm`
```
import {
  "math",
  "os"
}

const PI = 3.1415926

abstract class Shape {
  float area; # If you don't specify a starting value, it'll default to its zero value
  func calcArea() float;
}

class Rectangle : Shape {
  float length;
  float height;

  # Constructors do not have a `func` nor a return type
  Rectangle(float l, float h) {
    this.length = l;
    this.height = h;
    this.area = this.CalcArea();
  }

  func CalcArea() float {
    return this.length * this.height;
  }
}

class Square : Rectangle {
  Square(float side) {
    this.length = side;
    this.height = side;
    this.area = this.CalcArea();
  }

  # You can have multiple constructors with different structures
  Square(float l, float h) {
    parent(l, h); # Reference parent class using the `parent` keyword. To call its constructor, call it like a function
  }
}

class Circle : Shape {
  float radius;

  Circle(float r) {
    this.radius = r;
    this.area = this.CalcArea();
  }

  func CalcArea() float {
    return PI * this.radius^2;
  }
}

func main() int {
  var choice = math.Random(0:2); # <int>:<int>[:<int>] indicates a range, output as a list. 
  # Range format is <start>:<stop>[:<step>]. Ranges are INCLUSIVE

  # Declare something as a base time, and reassign it to a child type later on
  Shape shape;
  if choice == 0 {
    shape = Rectangle(1, 1);
  } elif choice == 1 {
    shape = Square(2);
  } elif (choice == 2) { # Optional parentheses in boolean expressions
    shape = Circle(2.5);
  } else {
    shape = null; # All types can optionally be null
  }

  if !shape {
    os.PrintErr("No shape defined!");
    return 1;
  }

  print("Shape: {type(shape)} -- Area: {shape.area}");

  return 0;
}
```
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

(Do we need to 'reserve' an octothorpe, or do we just entirely yeet the comments during Step 1/lexing and then not have to worry?)

### Keywords

- `if`
- `else`
- `elif` (?)
- `class`
- `return`
- `func`
- `defer`
- `const`
- `static`
- `private` (?)
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

- `list`
- `dict`

(if differentiating between `int`/`float` and `str`/`char` becomes too difficult, I'll collapse them into `num` and `str`, where a `char` is just a `str` with length 1)

## References

- [Compiler in Python](https://medium.com/@marcelogdeandrade/writing-your-own-programming-language-and-compiler-with-python-a468970ae6df)
- [TeenyTiny Compiler (also python)](https://austinhenley.com/blog/teenytinycompiler1.html)
- [Pinecone Development](https://www.freecodecamp.org/news/the-programming-language-pipeline-91d3f449c919/)
- [Crafting Interpreters Book](https://craftinginterpreters.com/)
