// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package main

import (
	"testing"

	"github.com/tabby-lang/tc/lexer"
	"github.com/tabby-lang/tc/token"
)

func TestToken(t *testing.T) {
	type Test struct {
		expectedType    token.Type
		expectedLiteral string
	}
	const input = `
			/* comment should not be scanned */
			five := "test"
			ten := 10
			add := fn(x, y) {
				x + y
			}

			result := add(five, ten)  
			5 < 10 > 5

			if (5 < 10) {
				return true
			} else {
				return false
			}

			10 == 10
			func add(x int
			`

	tests := []Test{
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("decl"), ":="},
		{token.TokMap.Type("string_literal"), "\"test\""},
		{token.TokMap.Type("ident"), "ten"},
		{token.TokMap.Type("decl"), ":="},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("ident"), "add"},
		{token.TokMap.Type("decl"), ":="},
		{token.TokMap.Type("ident"), "fn"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("comma"), ","},
		{token.TokMap.Type("ident"), "y"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("add"), "+"},
		{token.TokMap.Type("ident"), "y"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("ident"), "result"},
		{token.TokMap.Type("decl"), ":="},
		{token.TokMap.Type("ident"), "add"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("ident"), "five"},
		{token.TokMap.Type("comma"), ","},
		{token.TokMap.Type("ident"), "ten"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("int"), "5"},
		{token.TokMap.Type("lt"), "<"},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("gt"), ">"},
		{token.TokMap.Type("int"), "5"},
		{token.TokMap.Type("if"), "if"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("int"), "5"},
		{token.TokMap.Type("lt"), "<"},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("rparen"), ")"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("return"), "return"},
		{token.TokMap.Type("true"), "true"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("else"), "else"},
		{token.TokMap.Type("lbrace"), "{"},
		{token.TokMap.Type("return"), "return"},
		{token.TokMap.Type("false"), "false"},
		{token.TokMap.Type("rbrace"), "}"},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("eq"), "=="},
		{token.TokMap.Type("int"), "10"},
		{token.TokMap.Type("func"), "func"},
		{token.TokMap.Type("ident"), "add"},
		{token.TokMap.Type("lparen"), "("},
		{token.TokMap.Type("ident"), "x"},
		{token.TokMap.Type("ident"), "int"},
		{token.TokMap.Type("$"), ""}, // end token
	}

	l := lexer.NewLexer([]byte(input))
	for i, tt := range tests {
		tok := l.Scan()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected='%s', got='%s' at line %d, column %d",
				i, token.TokMap.Id(tt.expectedType), token.TokMap.Id(tok.Type), tok.Pos.Line, tok.Pos.Column)
		}

		if string(tok.Lit) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected='%q', got='%q' at line %d, column %d",
				i, tt.expectedLiteral, string(tok.Lit), tok.Pos.Line, tok.Pos.Column)
		}
	}
}
