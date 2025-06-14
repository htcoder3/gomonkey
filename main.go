package main

import (
	"fmt"
	"monkey/ast"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
	"os/user"
)

func testLetStatement(s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		// t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		// t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		// t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		// t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
		// 	name, letStmt.TokenLiteral())
		return false
	}

	return true
}

func runRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runMytest() {
	input := `len("aaa");`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := evaluator.Eval(program, env)
	fmt.Println(result.Inspect())

	// // checkParserErrors(t, p)

	// actual := program.String()
	// fmt.Println(actual)
}

func main() {

	runRepl()
	//

	// runMytest()
}
