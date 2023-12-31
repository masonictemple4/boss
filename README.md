# BOSS - An interpreted language built using go.
This repo follows along with the book [Writing An Interpreter](https://interpreterbook.com/) In Go by [Thorsten Ball](https://thorstenball.com/)

###### TODO:
* Lexer
* Parser
* AST
* Internal object system
* Evaluator

###### Notes from the book:
Programs in BOSS are series of statements.

**A let statement:** `let <identifier> = <expression>;`

**An example of a valid program witten in BOSS**:
```
let x = 10;
let y = 10;

let add fn(a, b) {
  return a + b;
};
```

Because our host language is Go, we benefit from it's built-in GC,
adding our own would require us to actively avoid and disable the Go
GC which is already a massive task ontop of managing our own memory. 

**A return statement:** `return <expression>;`

###### Challenges
- [ ] Implement a different `evaluation` strategy.


