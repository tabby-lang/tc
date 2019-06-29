// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tabby-lang/tc/ast"
	"github.com/tabby-lang/tc/checker"
	"github.com/tabby-lang/tc/gen"
	"github.com/tabby-lang/tc/lexer"
	"github.com/tabby-lang/tc/message"
	"github.com/tabby-lang/tc/parser"
)

func check(err error) {
	if err != nil {
		message.Error(err)
	}
}

func Parse(input string) *ast.Program {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	node, err := p.Parse(l)
	check(err)
	program, _ := node.(*ast.Program)
	return program
}

func TypeCheck(program *ast.Program) {
	err := checker.Checker(program)
	check(err)
}

func Compile(code bytes.Buffer) string {
	f, err := os.Create("./build/" + "main" + ".cpp")
	check(err)
	defer f.Close()
	f.Write(code.Bytes())

	var out bytes.Buffer
	cmd1 := exec.Command("g++", "-std=c++11", "-o", "main", "./build/"+"main.cpp", "-w", "./build/Builtins.cpp")
	cmd1.Stderr = &out
	err = cmd1.Run()

	// TODO: Make sure this only catches errors, not warnings.
	if len(out.String()) != 0 {
		message.Error(fmt.Sprintf("error: %s", out.String()))
	}

	cmd := exec.Command("./main.exe")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err = cmd.Run()

	if err != nil {
		message.Error(err.Error())
	}

	return outb.String()
}

func main() {
	if len(os.Args) < 2 {
		message.Error("No valid input file given!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	input, err := ioutil.ReadFile(absPath)
	check(err)

	program := Parse(string(input))
	TypeCheck(program)
	code := gen.GenWrapper(program)

	/** TODO: Refactor into "tc run" command.

	output := Compile(code)
	fmt.Println(output)

	*/

	Compile(code)
	message.Info("Code successfully compiled 😸")
}
