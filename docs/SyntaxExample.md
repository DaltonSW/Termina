# Syntax Example

Disclaimer:
- I wrote this out just kinda freeform to see what the syntax would end up like if I just wrote some code snippets how I wanted. Some of the ideas exhibited here might be stupid, or might be setting me up for a really hard time in the future

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
