// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/tabby-lang/tc/ast"
)

func TestAST(t *testing.T) {
	const input = `
			/* comment should not be scanned */
			func add() Int {
				return x + y
			}

			five := "test"
			ten := 10
			ten = 4

			result := 4  
			5 <= 10

			if (5 == 10) {
				print()
			} else {
				print()
			}

			10 == 10
			`

	out := &ast.Program{
		Functions: []ast.Statement{
			ast.FunctionStatement{Name: "add", Return: "Int", Parameters: []ast.FormalArg{}, Body: &ast.BlockStatement{
				Statements: []ast.Statement{ast.ReturnStatement{ReturnValue: ast.InfixExpression{Left: ast.StringLiteral{Value: "x"}, Operator: "+", Right: ast.StringLiteral{Value: "y"}}}}}}},
		Statements: []ast.Statement{
			ast.InitStatement{Expr: ast.StringLiteral{Value: "\"test\""}, Location: "five"},
			ast.InitStatement{Expr: ast.StringLiteral{Value: "10"}, Location: "ten"},
			ast.AssignStatement{Left: ast.Identifier{Value: "ten"}, Right: ast.IntegerLiteral{Value: "4"}},
			ast.InitStatement{Expr: ast.IntegerLiteral{Value: "4"}, Location: "result"},
			ast.ExpressionStatement{Expression: ast.InfixExpression{Left: ast.IntegerLiteral{Value: "5"}, Operator: "<=", Right: ast.IntegerLiteral{Value: "10"}}},
			ast.IfStatement{Condition: ast.InfixExpression{Left: ast.IntegerLiteral{Value: "5"}, Operator: "==", Right: ast.IntegerLiteral{Value: "10"}},
				Block:       &ast.BlockStatement{Statements: []ast.Statement{ast.ExpressionStatement{Expression: ast.FunctionCall{Name: "print", Args: []ast.Expression{}}}}},
				Alternative: &ast.BlockStatement{Statements: []ast.Statement{ast.ExpressionStatement{Expression: ast.FunctionCall{Name: "print", Args: []ast.Expression{}}}}}},
			ast.ExpressionStatement{Expression: ast.InfixExpression{Left: ast.IntegerLiteral{Value: "10"}, Operator: "==", Right: ast.IntegerLiteral{Value: "10"}}}}}

	program := Parse(input)

	js, _ := json.MarshalIndent(program, "", "    ")
	jsOut, _ := json.MarshalIndent(out, "", "    ")

	if !reflect.DeepEqual(js, jsOut) {
		fmt.Printf("\n%s\n", js)
		fmt.Println("****************************")
		fmt.Printf("\n%s\n", jsOut)

		t.Fatalf("Wrong AST")
	}

}
