// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package main

import (
	"testing"

	"github.com/tabby-lang/tc/ast"
	"github.com/tabby-lang/tc/checker"
	"github.com/tabby-lang/tc/lexer"
	"github.com/tabby-lang/tc/parser"
)

type Test struct {
	src     string
	Success bool
}

func TestOperations(t *testing.T) {
	tests := []Test{
		{`5 + 5`, true},
		{`5 + "5"`, false},
		{`5 * 5`, true},
		{`5 * 5 + 9`, true},
		{`5 < 10`, true},
		{`true and true`, true},
		{`4 and 2`, false},
		{`true or false`, true}}

	runTests(tests, t)
}

func TestIdents(t *testing.T) {
	tests := []Test{
		{`x := 5`, true},
		{`
			x := 6
			x := 8`, false},
		{
			`x := 5
			x = "hello"`, false},
		{
			`y := "hey"
			y = "cool"`, true},
		{
			`x := "hello "
			y := "world"
			z := x + y`, true},
		{
			`x = 5`, false}}

	runTests(tests, t)
}

func TestFunctions(t *testing.T) {
	tests := []Test{
		{
			`func add(x Int, y Int) Int {
				return x + y
			}

			a := add(1, 3)`, true},
		{
			`func add(x Int, y Int) Int {
				return x + y
			}

			z := add("test", 3)`, false},
		{
			`func one() Int {
				return "test"
			}`, false},
	}

	runTests(tests, t)
}

func runTests(tests []Test, t *testing.T) {
	for i, test := range tests {
		err := stringToChecker(test.src)

		if err != nil && !test.Success {
			continue
		} else if err != nil {
			t.Fatalf("test %d fail: "+err.Error(), i)
		} else if !test.Success {
			t.Fatalf("test %d should have failed", i)
		}
	}
}

func stringToChecker(input string) error {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	res, err := p.Parse(l)
	if err != nil {
		return err
	}

	program, _ := res.(*ast.Program)
	err = checker.Checker(program)
	if err != nil {
		return err
	}
	return nil
}
